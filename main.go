package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MovieNightBot struct {
	Db        *gorm.DB
	Bot       *tgbotapi.BotAPI
	OmdbToken string
}

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

type Response map[string]interface{}

func initializeDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		log.Panicf("DB Error: %s", err)
	}

	db.AutoMigrate(&Movie{}, &Vote{})

	return db
}

func initializeTgBot() *tgbotapi.BotAPI {
	botApiToken := os.Getenv("TG_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botApiToken)
	if err != nil {
		log.Panicf("Telegram API Error: %s", err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func handleUpdates(mnb *MovieNightBot) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := mnb.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "search":
			args := url.QueryEscape(update.Message.CommandArguments())
			url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s", mnb.OmdbToken, args)
			log.Printf("URL: %s", url)

			resp, _ := http.Get(url)
			defer resp.Body.Close()

			jsonBody, _ := io.ReadAll(resp.Body)

			var body Response
			json.Unmarshal(jsonBody, &body)

			statusCode := resp.StatusCode
			log.Printf("Status Code: %d", statusCode)

			if statusCode == 200 && body["Response"] == "True" {
				log.Printf("Entered Loop")
				var searchResults []map[string]interface{}
				json.Unmarshal([]byte(body["Search"].(string)), &searchResults)
				log.Printf("Search Results: %s", searchResults)
				for _, v := range searchResults {
					log.Printf("Found Result: %s", v["imdbID"].([]interface{}))
				}
			} else {
				log.Printf("Didn't enter loop. [Status Code: %d] [Response Value: %s]", statusCode, body["Response"])
			}
		}
	}
}

func main() {
	db := initializeDb()

	omdbToken := os.Getenv("OMDB_TOKEN")

	bot := initializeTgBot()

	mnb := &MovieNightBot{
		Db:        db,
		Bot:       bot,
		OmdbToken: omdbToken,
	}

	handleUpdates(mnb)
}
