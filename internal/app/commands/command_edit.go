package commands

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tg-bot/internal/service/product"
	"strconv"
	"strings"
)

func (commander *Commander) Edit(message *tgbotapi.Message) {
	args := message.CommandArguments()

	params := strings.Split(args, argDelimiter)
	if len(params) != 3 {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Required 3 fields, got %d", len(params))),
		)
		return
	}

	id, err := strconv.Atoi(params[0])
	if err != nil {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Id should be a positive integer or zero, got %s", params[0])),
		)
		return
	}

	price, err := strconv.Atoi(params[2])
	if err != nil {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Price should be a positive integer or zero, got %s", params[2])),
		)
		return
	}

	err = commander.productService.Update(uint(id), product.NewProduct(params[1], price))
	if err != nil {
		commander.HandleError(message.Chat.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("You have update a product with id %d", id),
	)
	commander.bot.Send(msg)
}
