package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// 问题描述
type QuestionDescription struct {
	Data Data `json:"data"`
}

type Data struct {
	IsCurrentUserAuthenticated bool       `json:"isCurrentUserAuthenticated"`
	UserStatus                 UserStatus `json:"userStatus"`
	Question                   Question   `json:"question"`
	SubscribeURL               string     `json:"subscribeUrl"`
	LoginURL                   string     `json:"loginUrl"`
}

type Question struct {
	QuestionID            string        `json:"questionId"`
	QuestionFrontendID    string        `json:"questionFrontendId"`
	QuestionTitle         string        `json:"questionTitle"`
	TranslatedTitle       string        `json:"translatedTitle"`
	QuestionTitleSlug     string        `json:"questionTitleSlug"`
	Content               string        `json:"content"`
	TranslatedContent     string        `json:"translatedContent"`
	Difficulty            string        `json:"difficulty"`
	EditorType            string        `json:"editorType"`
	Stats                 string        `json:"stats"`
	AllowDiscuss          bool          `json:"allowDiscuss"`
	Contributors          []interface{} `json:"contributors"`
	SimilarQuestions      string        `json:"similarQuestions"`
	MysqlSchemas          []interface{} `json:"mysqlSchemas"`
	RandomQuestionURL     string        `json:"randomQuestionUrl"`
	SessionID             string        `json:"sessionId"`
	CategoryTitle         string        `json:"categoryTitle"`
	SubmitURL             string        `json:"submitUrl"`
	InterpretURL          string        `json:"interpretUrl"`
	CodeDefinition        string        `json:"codeDefinition"`
	SampleTestCase        string        `json:"sampleTestCase"`
	EnableTestMode        bool          `json:"enableTestMode"`
	MetaData              string        `json:"metaData"`
	LangToValidPlayground string        `json:"langToValidPlayground"`
	EnableRunCode         bool          `json:"enableRunCode"`
	EnableSubmit          bool          `json:"enableSubmit"`
	JudgerAvailable       bool          `json:"judgerAvailable"`
	InfoVerified          bool          `json:"infoVerified"`
	EnvInfo               string        `json:"envInfo"`
	URLManager            string        `json:"urlManager"`
	Article               string        `json:"article"`
	QuestionDetailURL     string        `json:"questionDetailUrl"`
	LibraryURL            interface{}   `json:"libraryUrl"`
	TopicTags             []TopicTag    `json:"topicTags"`
	Typename              string        `json:"__typename"`
}

type TopicTag struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	TranslatedName string `json:"translatedName"`
	Typename       string `json:"__typename"`
}

type UserStatus struct {
	IsPremium bool   `json:"isPremium"`
	Typename  string `json:"__typename"`
}

// 代码定义

type Code []Element

type Element struct {
	Value       string `json:"value"`
	Text        string `json:"text"`
	DefaultCode string `json:"defaultCode"`
}

func GetQuestionDetail(client http.Client, title string, lang string, tpl string) (string, *Element, error) {
	url := "https://leetcode-cn.com/graphql"
	data := `{
    "operationName": "getQuestionDetail",
    "variables": {
        "titleSlug": "%s"
    },
    "query": "query getQuestionDetail($titleSlug: String!) {\n  isCurrentUserAuthenticated\n  userStatus {\n    isPremium\n    __typename\n  }\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    questionTitle\n    translatedTitle\n    questionTitleSlug\n    content\n    translatedContent\n    difficulty\n    editorType\n    stats\n    allowDiscuss\n    contributors {\n      username\n      profileUrl\n      __typename\n    }\n    similarQuestions\n    mysqlSchemas\n    randomQuestionUrl\n    sessionId\n    categoryTitle\n    submitUrl\n    interpretUrl\n    codeDefinition\n    sampleTestCase\n    enableTestMode\n    metaData\n    langToValidPlayground\n    enableRunCode\n    enableSubmit\n    judgerAvailable\n    infoVerified\n    envInfo\n    urlManager\n    article\n    questionDetailUrl\n    libraryUrl\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    __typename\n  }\n  subscribeUrl\n  loginUrl\n}\n"
	}`

	resp, err := client.Post(url, "application/json", strings.NewReader(fmt.Sprintf(data, title)))
	if err != nil {
		return "", nil, err

	}
	defer resp.Body.Close()

	// 反序列化body中题目的描述信息
	var d = QuestionDescription{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return "", nil, err
	}

	var content string
	switch lang {
	case "zh":
		content = d.Data.Question.TranslatedContent
	case "en":
		content = d.Data.Question.Content
	default:
		return "", nil, fmt.Errorf("Not support language: %s", lang)
	}

	// 反序列化返回结果中的代码定义
	code := Code{}
	err = json.Unmarshal([]byte(d.Data.Question.CodeDefinition), &code)
	if err != nil {
		return "", nil, err
	}
	var element Element
	for _, c := range code {
		if c.Text == tpl {
			element = c
			break
		}
	}
	return content, &element , nil
}
