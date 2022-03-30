package commands

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (commander *Commander) Delete(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Id should be a positive integer or zero, got %s", args)),
		)
		return
	}

	err = commander.productService.Remove(uint(id))
	if err != nil {
		commander.HandleError(message.Chat.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("You have deleted a product with id %d", id),
	)
	commander.bot.Send(msg)
}
