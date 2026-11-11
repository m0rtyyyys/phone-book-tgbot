package main

import (
	"context"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	var err error

	// Загрузка переменного окружения
	err = godotenv.Load()
	if err != nil {
		panic("Unable to open .env")
	}

	// Подключение к базе данных
	dsn := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic("Error while opening db")
	}
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Запуск бота
	token := os.Getenv("TG_BOT_KEY")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	// Цикл проверки на новые сообщения
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		res := make([]Profile, 0, 1)
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "name":
			var output string
			input := update.Message.CommandArguments()
			res, err = SearchByNameList(context.Background(), pool, input)
			if err != nil {
				msg.Text = fmt.Sprint("Что-то пошло не так: ", err)
			}
			output = fmt.Sprintf("Найдено совпадений по запросу \"%s\": %d\n", input, len(res))
			for _, persons := range res {
				output += fmt.Sprintln(persons)
			}
			msg.Text = output
		case "surname":
			var output string
			input := update.Message.CommandArguments()
			res, err = SearchBySurnameList(context.Background(), pool, input)
			if err != nil {
				msg.Text = fmt.Sprint("Что-то пошло не так: ", err)
			}
			output = fmt.Sprintf("Найдено совпадений по запросу \"%s\": %d\n", input, len(res))
			for _, persons := range res {
				output += fmt.Sprintln(persons)
			}
			msg.Text = output
		case "city":
			var output string
			input := update.Message.CommandArguments()
			res, err = SearchByCityList(context.Background(), pool, input)
			if err != nil {
				msg.Text = fmt.Sprint("Что-то пошло не так: ", err)
			}
			output = fmt.Sprintf("Найдено совпадений по запросу \"%s\": %d\n", input, len(res))
			for _, persons := range res {
				output += fmt.Sprintln(persons)
			}
			msg.Text = output
		case "help":
			msg.Text = fmt.Sprint("/name <текст> - поиск контакта по имени\n/surname <текст> - поиск контакта по фамилии")
		case "open":
			msg.Text = "nothing"
			msg.ReplyMarkup = numericKeyboard
		default:
			msg.Text = "Я не знаю такой комманды\n/help - для помощи"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}

	//res := make([]Profile, 0, 1)
	// res, err = SearchByNameList(context.Background(), pool, "иван")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("По именам\nНайдено совпадений: ", len(res))
	// for _, val := range res {
	// 	fmt.Println(val)
	// }
	// res, err = SearchBySurnameList(context.Background(), pool, "смир")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("По фамилиям\nНайдено совпадений: ", len(res))
	// for _, val := range res {
	// 	fmt.Println(val)
	// }
}
