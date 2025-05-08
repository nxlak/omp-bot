package profile

import (
	"encoding/json"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const limit uint64 = 5

func (c *UserProfileCommander) List(position uint64, inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products, err := c.profileService.List(position, limit)
	if err != nil {
		log.Printf("error - %v", err)
		return
	}

	for _, p := range products {
		outputMsgText += "ID: "
		outputMsgText += strconv.FormatUint(p.ID, 10)
		outputMsgText += " "
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 5,
	})

	callbackPath := path.CallbackPath{
		Domain:       "user",
		Subdomain:    "profile",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, errMsg := c.bot.Send(msg)
	if errMsg != nil {
		log.Printf("UserProfileCommander.List: error sending reply message to chat - %v", err)
	}
}
