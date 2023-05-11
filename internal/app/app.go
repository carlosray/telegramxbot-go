package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/op/go-logging"
	"telegramxbot/internal/cache"
	"telegramxbot/internal/config"
	"telegramxbot/internal/handler"
	ilog "telegramxbot/internal/logging"
	"telegramxbot/internal/model"
	"time"
)

var log = logging.MustGetLogger("xbot")

func Run() {
	start := time.Now()

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

	state := initAppState(handlers)

	log.Infof("Application started for %s. State: %v", state.StartedAt().Sub(start), state)

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		u := &update
		cache.Chats.AddChatFromUpdate(u)
		for name, h := range handlers {
			log.Debugf("Processing update #%d with handler \"%s\"", u.UpdateID, name)
			handled, err := (*h).Handle(bot, u)
			log.Debugf("Handled update #%d with handler \"%s\": %t", u.UpdateID, name, handled)
			if err != nil {
				log.Errorf("Error handling update #%d by \"%s\"", u.UpdateID, name, err)
			}
			if cfg.HandlePolicy == model.FIRST && handled {
				break
			}
		}
	}
}

func initAppState(handlers map[string]*handler.Handler) *cache.AppState {
	handlersNames := make([]string, 0, len(handlers))
	for k := range handlers {
		handlersNames = append(handlersNames, k)
	}

	state, err := cache.Initialize(time.Now(), handlersNames)
	if err != nil {
		log.Errorf("%v", err)
	}

	return state
}
