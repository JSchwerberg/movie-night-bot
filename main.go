package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title  string
	ImdbId string
}

type Vote struct {
	gorm.Model
	MovieID        int
	Movie          Movie
	TelegramUserId int
}

func handleUpdate(update *tgbotapi.Update) {
	log.Printf("%s %s: %s", update.Message.From.UserName, update.Message.Chat.Title, update.Message.Text)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		log.Panicf("DB Error: %s", err)
	}

	db.AutoMigrate(&Movie{}, &Vote{})

	botApiToken := os.Getenv("TG_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botApiToken)
	if err != nil {
		log.Panicf("Telegram API Error: %s", err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	omdbApiToken := os.Getenv("OMDB_API_TOKEN")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handleUpdate(&update)
	}
}
