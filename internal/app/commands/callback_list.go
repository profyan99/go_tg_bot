package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tg-bot/internal/app/pagination"
	"go-tg-bot/internal/app/path"
)

func (commander *Commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := pagination.CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		commander.HandleError(
			callback.Message.Chat.ID,
			err,
		)
		return
	}

	msg, err := commander.paginationService.BuildListMessage(
		callback.Message.Chat.ID,
		parsedData,
	)
	if err != nil {
		commander.HandleError(
			callback.Message.Chat.ID,
			err,
		)
		return
	}

	cb := tgbotapi.NewCallback(callback.ID, "")
	if _, err := commander.bot.Request(cb); err != nil {
		panic(err)
	}

	commander.bot.Send(msg)
}
