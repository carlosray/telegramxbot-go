package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegramxbot/internal/config"
)

type Handler interface {
	// Handle updates and return "true" if update processed to user
	Handle(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (bool, error)
	// Setup handler with properties from config
	Setup(props map[string]interface{})
}

var allHandlers = map[string]Handler{
	"hue": &HueHandler{},
}

func GetAllHandlers(cfg *config.Config) []*Handler {
	res := make([]*Handler, 0, len(allHandlers))
	for _, h := range cfg.Handlers {
		if handler := allHandlers[h.Name]; handler != nil {
			handler.Setup(h.Properties)
			res = append(res, &handler)
		}
	}
	return res
}
