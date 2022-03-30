package commands

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (commander *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Id should be a positive integer or zero, got %s", args)),
		)
		return
	}

	item, err := commander.productService.Get(uint(id))
	if err != nil {
		commander.HandleError(message.Chat.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Successfully parsed args: %s", item.String()))

	commander.bot.Send(msg)
}
