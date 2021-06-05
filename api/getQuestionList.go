package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type QuestionList struct {
	UserName        string           `json:"user_name"`
	NumSolved       int64            `json:"num_solved"`
	NumTotal        int64            `json:"num_total"`
	ACEasy          int64            `json:"ac_easy"`
	ACMedium        int64            `json:"ac_medium"`
	ACHard          int64            `json:"ac_hard"`
	StatStatusPairs []StatStatusPair `json:"stat_status_pairs"`
	FrequencyHigh   int64            `json:"frequency_high"`
	FrequencyMid    int64            `json:"frequency_mid"`
	CategorySlug    string           `json:"category_slug"`
}

type StatStatusPair struct {
	Stat       Stat        `json:"stat"`
	Status     interface{} `json:"status"`
	Difficulty Difficulty  `json:"difficulty"`
	PaidOnly   bool        `json:"paid_only"`
	IsFavor    bool        `json:"is_favor"`
	Frequency  int64       `json:"frequency"`
	Progress   int64       `json:"progress"`
}

type Difficulty struct {
	Level int64 `json:"level"`
}

type Stat struct {
	QuestionID          int64  `json:"question_id"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	TotalAcs            int64  `json:"total_acs"`
	TotalSubmitted      int64  `json:"total_submitted"`
	TotalColumnArticles int64  `json:"total_column_articles"`
	FrontendQuestionID  string `json:"frontend_question_id"`
	IsNewQuestion       bool   `json:"is_new_question"`
}

func GetQuestionList(client http.Client) (*QuestionList, error) {
	resp, err := client.Get("https://leetcode-cn.com/api/problems/all/")
	if err != nil {
		log.Println("send http request err: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	var res QuestionList
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Println("json decode err: ", err)
		return nil, err
	}
	return &res, nil
}
