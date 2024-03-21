package test

import (
	"chatbot/pkg"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getToken() (string, error) {
	err := godotenv.Load("../.env")

	if err != nil {
		return "", err
	}

	token := os.Getenv("OPENAI_KEY")

	return token, nil

}

func getUrl() (string, error) {
	err := godotenv.Load("../.env")

	if err != nil {
		return "", err
	}

	token := os.Getenv("OPENAI_HOST")

	return token, nil

}

func TestCompletion() {
	key, err := getToken()

	if err != nil {
		panic(err)
	}

	url, err := getUrl()

	if err != nil {
		panic(err)
	}

	chat := pkg.NewChatGPT(key, url, "gpt-3.5-turbo", 100)

	messages := []map[string]string{
		{"role": "system",
			"content": "Voce é um chabot do telegram responda as perguntas todas em caixa alta e começe com uma gargalhada sinistra tipo: AHAHAHAHAHAHHAHAHA"}, {
			"role":    "user",
			"content": "2 + 2",
		}}

	str, err := chat.Completion(messages)

	if err != nil {
		panic(err)
	}

	if str == "" {
		fmt.Println("string is empty")
	} else {
		fmt.Println(str)
	}

}
