package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (commander *Commander) List(message *tgbotapi.Message) {
	text := ""
	products := commander.productService.List()

	for _, productItem := range products {
		text += productItem.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "list_10"),
		),
	)

	commander.bot.Send(msg)
}
