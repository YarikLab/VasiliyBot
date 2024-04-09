package logger

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	HelloWord = "Привет! Я Вася, твой персональный помощник при учебе и подготовке к экзаменам.\n\n Я нахожусь в стадии разработки и на данный момент у тебя доступны 2 кнопки, попробуй испытать их!\n"
	PDFWord   = "Выбери учебник:"
	Ficha     = "Данная фича в разработке"
	Algebra   = tgbotapi.FilePath("./AlimovAlgebra10.pdf")
	SS        = "Выбери учебник"
	SS1       = tgbotapi.FilePath("./10_1_Углубленный Боголюбов.pdf")
	SS2       = tgbotapi.FilePath("./10_2_Углубленный Боголюбов.pdf")
	Day       = "Выбери день:"
	Mondey    = "1. Разговоры о важном\n2. Русский язык\n3. Химия\n4. Физика\n5. Английский язык\n6. Кубановедение\n7. Математика"
	Tuesday   = "1. Физическая культура\n2. Математика\n3. Биология\n4. Литература\n5. Математика\n6. Физика\n7. Английский язык"
	Wensday   = "1. Математика\n2. История\n3. Индивидуальный проект\n4. Русский язык \n5. Индивидуальный проект\n6. Информатика и ИКТ"
	Thrusday  = "1. Литература\n2. Актуальные вопросы обществознания\n3. Математика\n4. Психология и выбор профессии\n5. География\n6. История \n7. Общество"
	Friday    = "1. Литература\n2. Физическая культура\n3. История \n4. Общество\n5. Общество"
	Saturday  = "1. Общество\n2. Английский язык\n3. Физическая культура\n4. История\n5. Актуальные вопросы\n6. ОБЖ"
	Up        = "Идет загрузка, просьба не отправлять команды"
	CTOne     = "На данном этапе разработки ты можешь задать вопрос преподователю истории и общетсва по этой ссылке: \n\nhttp://t.me/anonimnyye_voprosy_bot?start=76VAub"
	GDZM      = "Только не говори учителям про эту функцию🤫\n\nhttps://gdz.ru/class-10/algebra/alimov-15/"
)

// главное меню
var StartMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("📔Учебный материал", "SendPDF"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"🤳Связь с преподавателем", "CT"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Расписание", "Uroki"),
	),
)

var GDZMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ГДЗ по математике", "GDZM"),
	),
)

// меню выбора учебников
var SendPDFMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Алгебра Алимов 10", "SendAlgebra"),
	),
)

// меню распиания
var RMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Понидельник", "M"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Вторник", "T"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Среда", "W"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Четверг", "Thru"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Пятница", "F"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Суббота", "S"),
	),
)

// инициализируем комнады
func Commands(Update tgbotapi.Update) {
	Command := Update.Message.Command()
	switch Command {
	case "start":
		MSG := tgbotapi.NewMessage(Update.Message.Chat.ID, HelloWord)
		MSG.ReplyMarkup = StartMenu
		MSG.ParseMode = "Markdown"
		SendMessage(MSG)
	case "send":
		MSG := tgbotapi.NewMessage(Update.Message.Chat.ID, PDFWord)
		MSG.ReplyMarkup = SendPDFMenu
		MSG.ParseMode = "Markdown"
		SendMessage(MSG)
	case "gdz":
		MSG := tgbotapi.NewMessage(Update.Message.Chat.ID, "Это скрытая функция, не отображающаяся в списке команд. Выбери ссылку")
		MSG.ReplyMarkup = GDZMenu
		MSG.ParseMode = "Markdown"
		SendMessage(MSG)
	}

}

// инициализируем колбеки
func Callback(Update tgbotapi.Update) {
	data := Update.CallbackQuery.Data
	chatID := Update.CallbackQuery.From.ID
	switch data {
	case "SendPDF":
		MSG := tgbotapi.NewMessage(chatID, PDFWord)
		MSG.ReplyMarkup = SendPDFMenu
		MSG.ParseMode = "Markdown"
		SendMessage(MSG)
	case "SendAlgebra":
		MSG := tgbotapi.NewMessage(chatID, "Идет загрузка . . .")
		SendMessage(MSG)
		file := tgbotapi.NewDocument(chatID, Algebra)
		Bot.Send(file)
	case "CT":
		MSG := tgbotapi.NewMessage(chatID, CTOne)
		SendMessage(MSG)
	case "Uroki":
		MSG := tgbotapi.NewMessage(chatID, Day)
		MSG.ReplyMarkup = RMenu
		MSG.ParseMode = "Markdown"
		SendMessage(MSG)
	case "M":
		MSG := tgbotapi.NewMessage(chatID, Mondey)
		SendMessage(MSG)
	case "T":
		MSG := tgbotapi.NewMessage(chatID, Tuesday)
		SendMessage(MSG)
	case "W":
		MSG := tgbotapi.NewMessage(chatID, Wensday)
		SendMessage(MSG)
	case "Thru":
		MSG := tgbotapi.NewMessage(chatID, Thrusday)
		SendMessage(MSG)
	case "F":
		MSG := tgbotapi.NewMessage(chatID, Friday)
		SendMessage(MSG)
	case "S":
		MSG := tgbotapi.NewMessage(chatID, Saturday)
		SendMessage(MSG)
	case "GDZM":
		MSG := tgbotapi.NewMessage(chatID, GDZM)
		SendMessage(MSG)
	}
}

// функция отправки сообщений с обработкой ошибок
func SendMessage(msg tgbotapi.Chattable) {
	if _, err := Bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
