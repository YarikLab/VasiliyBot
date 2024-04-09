package logger

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var Bot *tgbotapi.BotAPI

func StartBot() {

	// загружаем токен
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// объявляем переменную, которая связывает телеграмм и код
	Bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Failed get tocen", err)
	}

	// прописываем дебаг
	Bot.Debug = true
	log.Printf("запустился бот %v", Bot.Self.UserName)

	// инициализируем конфиг обновлений
	U := tgbotapi.NewUpdate(0)
	U.Timeout = 60
	Updates := Bot.GetUpdatesChan(U)
	if err != nil {
		log.Fatal("Failed listening updates", err)
	}

	// цикл обновлений
	for Update := range Updates {
		if Update.CallbackQuery != nil {
			Callback(Update)
		} else if Update.Message.IsCommand() {
			Commands(Update)
		}
	}

}

