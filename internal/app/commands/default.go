package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (commander *Commander) DefaultBehavior(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Ты написал: %s", message.Text))

	commander.bot.Send(msg)
}
