package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatGPT struct {
	key string
	url string

	model string

	maxTokens int16
}

func NewChatGPT(key, url, model string, maxTokens int16) *ChatGPT {
	return &ChatGPT{
		key:       key,
		url:       url,
		model:     model,
		maxTokens: maxTokens,
	}
}

func (c *ChatGPT) Completion(message string) (string, error) {

	// creating the request body
	requestData := map[string]interface{}{
		"model":      c.model,
		"max_tokens": c.maxTokens,
		"messages": []interface{}{map[string]interface{}{
			"role":    "user",
			"content": message,
		}},
	}

	// marshaling the request

	requestBody, err := json.Marshal(requestData)

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+c.key)
	request.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var d map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&d)

	/*
		if i chose to user the json.Unmarshal
		i need to transform the response.body in slice of bytes
		using io.Readall()
	*/

	if err != nil {
		return "", err
	}

	result := fmt.Sprint(d["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"])

	return result, nil
}
