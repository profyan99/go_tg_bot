package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tg-bot/internal/app/path"
	"go-tg-bot/internal/service/product"
	"log"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (commander *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		commander.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		commander.handleMessage(update.Message)
	}
}

func (commander *Commander) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Commander.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.CallbackName {
	case "list":
		commander.CallbackList(callback, callbackPath)
	default:
		log.Printf("Commander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (commander *Commander) handleMessage(message *tgbotapi.Message) {
	if !message.IsCommand() {
		commander.showCommandFormat(message)

		return
	}

	switch message.Command() {
	case "help":
		commander.Help(message)
	case "list":
		commander.List(message)
	case "get":
		commander.Get(message)
	case "delete":
		commander.Delete(message)
	case "new":
		commander.New(message)
	case "edit":
		commander.Edit(message)
	default:
		log.Printf("User %s wrotes undefined command: %s", message.From.UserName, message.Text)
	}
}

func (commander *Commander) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}")

	commander.bot.Send(outputMsg)
}

func (commander *Commander) HandleError(chatID int64, err error) {
	log.Printf("An error has ocurred in Commander: %v", err)

	outputMsg := tgbotapi.NewMessage(chatID, err.Error())

	if _, err := commander.bot.Send(outputMsg); err != nil {
		log.Printf("An error has ocurred in HandleError while sending error to user: %v", err)
	}
}
