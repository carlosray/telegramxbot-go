package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"telegramxbot/internal/cache"
	"telegramxbot/internal/util"
	"time"
)

type StatusCommandHandler struct {
	Command
	environment string
	timeFormat  string
}

func (s *StatusCommandHandler) Setup(props map[string]interface{}) {
	s.cmd = props["command"].(string)
	s.environment = props["environment"].(string)
	s.timeFormat = time.RFC1123
}

const msgTpl = `Hello, @%s!
I'm Xbot.

Environment: %s
Chat id: %d
Active Handlers: %s
Started At: %s
Uptime: %s
`

func (s *StatusCommandHandler) Handle(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (bool, error) {
	var msg *tgbotapi.Message
	var ok bool
	if msg, ok = util.GetCommandMessage(update); !ok || msg.Command() != s.cmd {
		return false, nil
	}

	log.Infof("Processing 'status' command from chat %d", msg.Chat.ID)

	state, err := cache.GetAppState()
	if err != nil {
		return false, err
	}

	message := fmt.Sprintf(msgTpl,
		msg.From.UserName,
		s.environment,
		msg.Chat.ID,
		strings.Join(state.Handlers(), ","),
		state.StartedAt().Format(s.timeFormat),
		time.Now().Sub(state.StartedAt()),
	)

	msgToSend := tgbotapi.NewMessage(msg.Chat.ID, message)
	msgToSend.ReplyToMessageID = msg.MessageID

	if _, err := bot.Send(msgToSend); err != nil {
		return false, err
	}
	return true, nil
}
