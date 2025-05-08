package user

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/user/profile"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type UserCommander struct {
	bot              *tgbotapi.BotAPI
	profileCommander Commander
}

func NewUserCommander(
	bot *tgbotapi.BotAPI,
) *UserCommander {
	return &UserCommander{
		bot: bot,
		// profileCommander
		profileCommander: profile.NewUserProfileCommander(bot),
	}
}

func (c *UserCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "profile":
		c.profileCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("UserCommander.HandleCallback: unknown profile - %s", callbackPath.Subdomain)
	}
}

func (c *UserCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "profile":
		c.profileCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown profile - %s", commandPath.Subdomain)
	}
}
