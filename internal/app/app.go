package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/op/go-logging"
	"telegramxbot/internal/cache"
	"telegramxbot/internal/config"
	"telegramxbot/internal/handler"
	ilog "telegramxbot/internal/logging"
	"telegramxbot/internal/model"
)

var log = logging.MustGetLogger("xbot")

func Run() {
	cfg := config.LoadConfig()
	ilog.Configure(cfg)

	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Errorf("Error creating bot", err)
		panic(err)
	}

	bot.Debug = cfg.Bot.Debug

	log.Infof("Bot is %v", bot.Self)

	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to N seconds on each request for an update.
	updateConfig.Timeout = cfg.UpdateConfig.Timeout

	handlers := handler.GetAllHandlers(cfg)

	if handlers == nil || len(handlers) <= 0 {
		panic("setup at least one handler")
	}

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		cache.Chats.AddChatFromUpdate(&update)
		for _, h := range handlers {
			handled, err := (*h).Handle(bot, &update)
			if err != nil {
				log.Errorf("Error handling update", err)
			}
			if cfg.HandlePolicy == model.FIRST && handled {
				break
			}
		}
	}
}
