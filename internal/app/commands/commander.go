package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"telegram-bot/internal/service/product"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: service,
	}
}

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v\n", parsedData),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		command, ok := registeredCommands[update.Message.Command()]
		if ok {
			command(c, update.Message)
		} else {
			c.Default(update.Message)
		}
	}
}
