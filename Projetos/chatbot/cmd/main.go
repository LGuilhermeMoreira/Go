package main

import (
	"chatbot/test"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getToken() string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalln(err)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	return token
}

func main() {
	// // creating a new bot
	// bot, err := tgbotapi.NewBotAPI(getToken())

	// if err != nil {
	// 	log.Panic(err)
	// }

	// // set debug true
	// bot.Debug = true

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	// //create the message updater
	// u := tgbotapi.NewUpdate(0)
	// // setting the timeout message updater
	// u.Timeout = 60

	// updates := bot.GetUpdatesChan(u)

	// // searching for messages
	// for update := range updates {
	// 	if update.Message != nil { // if we got a message
	// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	// 		result, err := pkg.Completion(update.Message.Text)

	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		// creating a message
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
	// 		msg.ReplyToMessageID = update.Message.MessageID

	// 		// sending the message
	// 		bot.Send(msg)
	// 	}
	// }

	test.TestCompletion()
}
