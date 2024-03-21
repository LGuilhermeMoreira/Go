package pkg

import (
	"bytes"
	"chatbot/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChatGPT struct {
	key string
	url string

	model string

	maxTokens uint32
}

func NewChatGPT(key, url, model string, maxTokens uint32) *ChatGPT {
	return &ChatGPT{
		key:       key,
		url:       url,
		model:     model,
		maxTokens: maxTokens,
	}
}

func (c *ChatGPT) Completion(messages []map[string]string) (string, error) {

	var data models.Messages

	requestBody := map[string]interface{}{
		"model":      c.model,
		"max_tokens": c.maxTokens,
		"messages":   messages,
	}

	requestMarshalled, err := json.Marshal(requestBody)

	if err != nil {
		fmt.Printf("error in Completion marshal the requestbody: %v", err)
		return "", err
	}

	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestMarshalled))

	if err != nil {
		fmt.Printf("error in Completion create newRequest: %v", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+c.key)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Printf("error in Completion do http request: %v", err)
		return "", nil
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		fmt.Printf("error in Completion decode the response.body: %v", err)
		return "", err
	}

	fmt.Println(data)

	var result string

	if len(data.Choices) > 0 {
		result = data.Choices[0].Text
	} else {
		result = "Resposta vazia"
	}

	fmt.Println(result)

	return result, nil
}
