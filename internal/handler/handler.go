package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/op/go-logging"
	"telegramxbot/internal/config"
)

var log = logging.MustGetLogger("xbot")

type Handler interface {
	// Setup handler with properties from config
	Setup(props map[string]interface{})
	// Handle updates and return "true" if update processed to user
	Handle(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (bool, error)
}

type Command struct {
	cmd string
}

var allHandlers = map[string]Handler{
	"hue":    &HueHandler{},
	"status": &StatusCommandHandler{},
}

func GetAllHandlers(cfg *config.Config) map[string]*Handler {
	res := make(map[string]*Handler)
	for _, h := range cfg.Handlers {
		if handler, ok := allHandlers[h.Name]; ok && handler != nil {
			handler.Setup(h.Properties)
			res[h.Name] = &handler
			log.Infof("Handler \"%s\" of type %T configured", h.Name, handler)
		}
	}
	return res
}
