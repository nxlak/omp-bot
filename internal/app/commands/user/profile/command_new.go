package profile

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *UserProfileCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, errCr := c.profileService.Create(args)
	if errCr != nil {
		log.Printf("error creating new entity - %v", errCr)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Entity with id=%d created", idx),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.New: error sending reply message to chat - %v", err)
	}
}
