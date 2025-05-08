package profile

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *UserProfileCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	entityID, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	done, err := c.profileService.Remove(uint64(entityID))
	if !done {
		log.Printf("Entity with id %d doesn't exists: %v", entityID, err)
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
