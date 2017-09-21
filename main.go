package main

import (
	"log"
	"github.com/gopkg.in/telegram-bot-api-4.6"
	"fmt"
	"encoding/json"
	"os"
	//"go/token"
)

//Создание массива
type Config struct {
	Listbase []Listbase
}
type Listbase struct {
	Name  string `json:"name"`
	List1 List1  `json:"list1"`
	List2 List2  `json:"list2"`
	List3 List3  `json:"list3"`
}
type List1 struct {
	List_name   string `json:"list_name"`
	List_text   string `json:"list_text"`
	List_status string `json:"list_status"`
}
type List2 struct {
	List_name   string `json:"list_name"`
	List_text   string `json:"list_text"`
	List_status string `json:"list_status"`
}
type List3 struct {
	List_name   string `json:"list_name"`
	List_text   string `json:"list_text"`
	List_status string `json:"list_status"`
}

//здание config
func loadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, err
}

//p1:= Person{"Sam", 20, []string{}}

func main() {
	config, _ := loadConfiguration("config.json")
	//
	fmt.Println(config.Listbase[0].Name)

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
		if update.Message.Text == "/start" {
			var reply string
			if update.Message.From.UserName != "" {
				// В чат вошел новый пользователь
				// Поприветствуем его
				reply = fmt.Sprintf(`Привет @%s! Я тут слежу за порядком. Веди себя хорошо. Удинственная командакоторую я сейчас пытаюсь выполнить это /create, но и она не очень то и хорошо работает!`,
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
		if update.Message.Text == "/create" {
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
			if update.Message.Text != "" {

			}
			continue
		}
		if update.Message.Text == "name" {
			var reply string
			if update.Message.From.UserName != "" {
				reply = fmt.Sprintf(config.Listbase[0].Name, config.Listbase[0].List1.List_name)
			}
			if reply != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			}
			continue
		}
		if update.Message.Text == "dela" {
			var reply string
			if update.Message.From.UserName != "" {
				dela := config.Listbase[0].List1.List_name
				reply = fmt.Sprintf(dela, config.Listbase[0].List1.List_name)
			}
			if reply != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			}
			continue
		}
	}
}
