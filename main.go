package main

import (
	"log"
	"github.com/gopkg.in/telegram-bot-api-4.6"
	"fmt"
	"encoding/json"
	"os"


)
type Config struct {
	Database struct{
		Host string `json:"host"`
	}`json:"database"`
	Host string `json:"host"`
}
func loadConfiguration(filename string)(Config, error)  {
	var config Config
	configFile, err := os.Open(filename)
	defer  configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser:=json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config,err
}

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Print("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.Text=="/start"{
			var reply string
			if update.Message.From.UserName != "" {
				// В чат вошел новый пользователь
				// Поприветствуем его
				reply = fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо.`,
					update.Message.From.UserName)
			}
			if reply != "" {
				// Созадаем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				// и отправляем его
				bot.Send(msg)
			}
			continue
		}
		if update.Message.Text=="/create"{
			var reply string
			if update.Message.From.UserName != "" {
				// В чат вошел новый пользователь
				// Поприветствуем его
				reply = fmt.Sprintf(`@%s, напишите название заметки`,
					update.Message.From.UserName)
			}
			if reply != "" {
				// Созадаем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				// и отправляем его
				bot.Send(msg)
			}
			if update.Message.Text!="" {

			}
			continue
		}
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//bot.Send(msg)
	}
}

