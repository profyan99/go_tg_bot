package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (commander *Commander) Help(message *tgbotapi.Message) {
	commands := "/help - print list of commands\n" +
		"/list - get a list of entities\n" +
		"/get - get an entity\n" +
		"/delete - delete an existing entity\n" +
		"/new - create a new entity\n" +
		"/edit - edit an existing entity"

	msg := tgbotapi.NewMessage(message.Chat.ID, commands)

	commander.bot.Send(msg)
}
