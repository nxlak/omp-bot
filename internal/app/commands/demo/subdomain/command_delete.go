package subdomain

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	entityID, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	done := c.subdomainService.Delete(uint64(entityID))
	if !done {
		log.Printf("Entity with id %d doesn't exists", entityID)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Entity deleted",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Delete: error sending reply message to chat - %v", err)
	}
}
