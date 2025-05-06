package subdomain

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DemoSubdomainCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	parts := strings.SplitN(args, " ", 2)
	entityID, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	done := c.subdomainService.Edit(uint64(entityID), parts[1])
	if !done {
		log.Printf("Entity with id %d doesn't exists", entityID)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Entity witd id=%d edited", entityID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Edit: error sending reply message to chat - %v", err)
	}
}
