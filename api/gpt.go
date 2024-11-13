package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TongYiClient struct {
	apiKey string
}

type TongYiRsp struct {
	Output struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"output"`
	Usage struct {
		OutputTokens int `json:"output_tokens"`
		InputTokens  int `json:"input_tokens"`
	} `json:"usage"`
	RequestID string `json:"request_id"`
}

func NewTongYiClient(apiKey string) *TongYiClient {
	return &TongYiClient{
		apiKey: apiKey,
	}
}

func (c *TongYiClient) GenerateText(ctx context.Context, prompt string, history ...map[string]string) (*TongYiRsp, error) {
	// 构建请求数据
	data := map[string]interface{}{
		"model":      "qwen-turbo",
		"parameters": map[string]interface{}{},
		"input": map[string]interface{}{
			"prompt":  prompt,
			"history": history,
		},
	}

	// 转换为JSON格式
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 构建请求
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// 设置请求头部
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码是否为200
	if resp.StatusCode != http.StatusOK {
		var errorResponse struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		}

		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("API error: %s - %s", errorResponse.Code, errorResponse.Message)
	}

	// 读取响应数据
	response := &TongYiRsp{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.Output.Text)
	// 提取生成的文本并返回
	return response, nil
}
