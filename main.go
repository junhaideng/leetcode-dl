package main

import (
	"flag"
	"fmt"
	"leetcode/api"
	"leetcode/code"
	"leetcode/log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var client http.Client

var dir string
var template string
var lang string
var maxRetry int
var sleep uint
var write string // 日志是否写入文件中
var inc bool

func init() {
	rand.Seed(time.Now().Unix())
	flag.StringVar(&dir, "d", "leetcode", "directory to store all problems")
	flag.StringVar(&template, "t", "Go", "which program language to solve, [C, C++, Python, Python3, Java, etc]")
	flag.StringVar(&lang, "lang", "zh", "description language, support [zh, en]")
	flag.IntVar(&maxRetry, "r", 5, "get question detail retry times")
	flag.UintVar(&sleep, "s", 5, "sleep time to avoid http code 429")
	flag.StringVar(&write, "w", "", "write log to file")
	flag.BoolVar(&inc, "a", false, "if question exists, not download again")
}

func main() {
	flag.Parse()
	if write != "" {
		log.SetOutput(write)
	}
	// 首先获取到所有的问题列表
	resp, err := api.GetQuestionList(client)
	if err != nil {
		log.Errorf("获取题目列表失败: %s\n", err)
		return
	}
	err = os.Mkdir(dir, 0777)
	if err != nil {
		if !os.IsExist(err) {
			log.Errorf("创建文件夹[%s]失败: %s", dir, err)
			return
		}
	}

	for _, pair := range resp.StatStatusPairs {
		log.Infof("正在获取[%s]的问题详情描述\n", pair.Stat.QuestionTitle)
		if pair.PaidOnly {
			log.Warn("问题需要付费，跳过")
			continue
		}

		problemDir := filepath.Join(dir, pair.Stat.FrontendQuestionID)
		if inc {
			_, err := os.Stat(problemDir)
			if err == nil || os.IsExist(err) {
				log.Successf("题目文件夹已存在，跳过\n", pair.Stat.QuestionTitle, filepath.Join(problemDir))
				continue
			}
		}

		// 获取到每一个题目的具体信息
		desc, codeEle, err := api.GetQuestionDetail(client, pair.Stat.QuestionTitleSlug, lang, template)
		retries := 1

		for err != nil && retries <= maxRetry {
			// 随机 sleep 一段时间，避免发送太多请求了
			// hhhh，由此推出 leetcode 这部分 API 有个限流器
			// 也不知道是令牌桶还是令牌漏斗还是计数器还是滑动窗口呢 🤔
			time.Sleep(time.Second * time.Duration(rand.Intn(int(sleep))))

			log.Errorf("获取题目[%s]描述失败: %s\n", pair.Stat.QuestionTitle, err)
			log.Infof("进行第%d次重试\n", retries)

			desc, codeEle, err = api.GetQuestionDetail(client, pair.Stat.QuestionTitleSlug, lang, template)
			retries++
		}
		if err != nil {
			log.Errorf("获取题目描述失败，跳过题目 [id=%d, title=%s]\n", pair.Stat.QuestionID, pair.Stat.QuestionTitle)
			continue
		}

		//  -----------开始写入文件-----------------------
		err = os.MkdirAll(problemDir, 0777)
		if err != nil {
			if !os.IsExist(err) {
				log.Errorf("创建文件夹[%s]失败: %s", dir, err)
				return
			}
		}
		// 保存题目
		err = os.WriteFile(filepath.Join(problemDir, "README.md"), []byte(desc), 0666)
		if err != nil {
			log.Errorf("写入题目[%s]文件描述失败: %s\n", pair.Stat.QuestionTitle, err)
			continue
		}

		// 保存代码模板
		if codeEle != nil {
			err = os.WriteFile(filepath.Join(problemDir, fmt.Sprintf("%s%s", "main", code.GetTplSuffix(template))), []byte(codeEle.DefaultCode), 0666)
			if err != nil {
				log.Errorf("写入题目[%s]代码模板失败: %s", pair.Stat.QuestionTitle, err)
				continue
			}
			log.Successf("成功保存题目以及代码[%s]到目录[%s]\n", pair.Stat.QuestionTitle, filepath.Join(problemDir))
		}

	}
	log.Successf("全部题目下载完成，Enjoy!!")
}
