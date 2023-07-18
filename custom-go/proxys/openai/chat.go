package openai

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"encoding/json"
	"fmt"
	"github.com/artisancloud/openai"
	v1 "github.com/artisancloud/openai/api/v1"
)

func init() {
	plugins.AddProxyHook(chat, nil)
}

func chat(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	// 创建一个OpenAI API客户端
	config := openai.V1Config{
		OpenAPIKey: "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",
		ProxyURL:   "http://127.0.0.1:17890/",
	}
	client, err := openai.NewClient(&config)
	if err != nil {
		panic(err)
	}

	// 创建 ChatCompletion

	question := "你觉得小白的成绩怎么样?"
	req1 := v1.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []v1.Message{
			{
				Role:    "user",
				Content: question,
			},
		},
		Functions: []*v1.Function{
			{
				Name:        "getScore",
				Description: "通过学生的姓名，查询学生的成绩",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"name": map[string]string{
							"type":        "string",
							"description": "学生的姓名",
						},
					},
					"required": []string{"name"},
				},
			},
		},
		FunctionCall: "auto",
	}
	resp, err1 := client.V1.Chat.CreateChatCompletion(&req1)
	if err1 != nil {
		if e, ok := err1.(*v1.Error); ok {
			fmt.Printf("open ai error: %v", e.Detail.Message)
		}
		fmt.Printf("error: %v", err1)
		return nil, err1
	}

	// 输出结果
	message := resp.Choices[0].Message
	arg, _ := message.FunctionCall.(map[string]interface{})["arguments"]

	err = json.Unmarshal([]byte(arg.(string)), &arg)
	if err != nil {
		fmt.Printf("JSON序列化错误:%v", err)
	}

	name := arg.(map[string]interface{})["name"]
	fmt.Printf("name=%s", name)

	return body.Response, nil
}

func getScore(name string) string {
	// 用一个字典存储姓名和成绩对应关系
	scores := map[string]int{"小明": 90, "小红": 80, "小白": 59}
	score := map[string]interface{}{"name": name, "score": scores[name]}
	jsonStr, err := json.Marshal(score)
	if err != nil {
		fmt.Printf("JSON序列化错误:%v", err)
	}

	return string(jsonStr)
}
