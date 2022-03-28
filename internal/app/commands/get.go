package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (commander *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Print("Wrong args: ", args)
		return
	}

	item, err := commander.productService.Get(id)
	if err != nil {
		log.Print("Fail to get product: ", id, err)
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Successfully parsed args: %s", item.Title))

	commander.bot.Send(msg)
}
