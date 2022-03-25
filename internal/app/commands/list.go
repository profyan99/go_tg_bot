package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (commander *Commander) List(message *tgbotapi.Message) {
	text := ""
	products := commander.productService.List()

	for _, productItem := range products {
		text += productItem.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	commander.bot.Send(msg)
}
