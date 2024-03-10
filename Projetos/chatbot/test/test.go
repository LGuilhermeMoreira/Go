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

	chat := pkg.NewChatGPT(key, url, "babbage-002", 100)

	str, err := chat.Completion("2+2")

	if err != nil {
		panic(err)
	}

	fmt.Println(str)
}
