package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander *Commander) List(message *tgbotapi.Message) {
	text := ""
	products, err := commander.productService.List(0, 10)
	if err != nil {
		commander.HandleError(message.Chat.ID, err)
		return
	}

	for id, productItem := range products {
		text += fmt.Sprintf("[%d] %s - %d $\n", id, productItem.Name, productItem.Price)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "list_10"),
		),
	)

	commander.bot.Send(msg)
}
