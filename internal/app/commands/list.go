package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Here all the products\n\n"
	products := c.service.List()

	for _, p := range products {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
