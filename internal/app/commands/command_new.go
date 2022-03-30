package commands

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tg-bot/internal/service/product"
	"strconv"
	"strings"
)

func (commander *Commander) New(message *tgbotapi.Message) {
	args := message.CommandArguments()

	params := strings.Split(args, argDelimiter)
	if len(params) != 2 {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Required 2 fields, got %d", len(params))),
		)
		return
	}

	amount, err := strconv.Atoi(params[1])
	if err != nil {
		commander.HandleError(
			message.Chat.ID,
			errors.New(fmt.Sprintf("Amount should be a positive number, got %s", params[1])),
		)
		return
	}
	newProduct := product.NewProduct(params[0], amount)
	index := commander.productService.Add(newProduct)

	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		fmt.Sprintf("You have created a new Product %s with id %d", newProduct.Name, index),
	)
	commander.bot.Send(msg)
}
