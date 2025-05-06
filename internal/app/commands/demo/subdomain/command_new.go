package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	c.subdomainService.New(args)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"New entity created",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.New: error sending reply message to chat - %v", err)
	}
}
