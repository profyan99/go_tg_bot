package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (commander *Commander) Help(message *tgbotapi.Message) {
	commands := "/help - список всех команд\n" +
		"/create - создать ивент\n" +
		"/list - список текущих и созданных ивентов\n" +
		"/join - присоединиться к ивенту\n" +
		"/leave - покинуть ивент\n" +
		"/notify - оповестить учатников ивента"

	msg := tgbotapi.NewMessage(message.Chat.ID, commands)

	commander.bot.Send(msg)
}
