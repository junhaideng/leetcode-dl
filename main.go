package main

import (
	"flag"
	"fmt"
	"leetcode/api"
	"leetcode/log"
	"leetcode/code"
	"net/http"
	"os"
	"path/filepath"
)

var client http.Client

var dir string
var template string
var lang string

func init() {
	flag.StringVar(&dir, "d", "leetcode", "directory to store all problems")
	flag.StringVar(&template, "t", "Go", "which program language to solve, [C, C++, Python, Python3, Java, etc]")
	flag.StringVar(&lang, "lang", "zh", "description language, support [zh, en]")
}

func main() {
	flag.Parse()
	// 首先获取到所有的问题列表
	resp, err := api.GetQuestionList(client)
	if err != nil {
		log.Errorf("获取题目列表失败: %s\n", err)
		return
	}
	err = os.Mkdir(dir, 0666)
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

		// 获取到每一个题目的具体信息
		desc, codeEle, err := api.GetQuestionDetail(client, pair.Stat.QuestionTitleSlug, lang, template)
		if err != nil {
			log.Errorf("获取题目[%s]描述失败: %s\n", pair.Stat.QuestionTitle, err)
			continue
		}

		//  -----------开始写入文件-----------------------
		problemDir := filepath.Join(dir, pair.Stat.FrontendQuestionID)
		err = os.Mkdir(problemDir, 0666)
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
