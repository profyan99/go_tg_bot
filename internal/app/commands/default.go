package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (commander *Commander) DefaultBehavior(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Ты написал: %s", message.Text))

	commander.bot.Send(msg)
}

func (commander *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		args := strings.Split(update.CallbackQuery.Data, "_")
		page, _ := strconv.Atoi(args[1])

		text := fmt.Sprintf("Command: %s\nOffset: %d", args[0], page)
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
		commander.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		commander.Help(update.Message)
	case "list":
		commander.List(update.Message)
	case "get":
		commander.Get(update.Message)
	default:
		commander.DefaultBehavior(update.Message)
	}
}
