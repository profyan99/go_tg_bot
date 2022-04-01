package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tg-bot/internal/app/pagination"
)

func (commander *Commander) List(message *tgbotapi.Message) {
	msg, err := commander.paginationService.BuildListMessage(message.Chat.ID, pagination.CallbackListData{})
	if err != nil {
		commander.HandleError(message.Chat.ID, err)
	}

	commander.bot.Send(msg)
}
