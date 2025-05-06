package subdomain

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	entityID, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	entity, err := c.subdomainService.Get(uint64(entityID))
	if err != nil {
		log.Printf("failed to get entity with id %d: %v", entityID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Found entity with id=%v\nTitle: %s", entityID, entity.Title),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
