package cache

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ChatCache struct {
	private []int64
	group   []int64
}

var Chats = &ChatCache{}

func (c *ChatCache) AddChatFromUpdate(update *tgbotapi.Update) {
	if update != nil && update.Message != nil && update.Message.Chat != nil {
		c.AddChat(update.Message.Chat)
	}
}

func (c *ChatCache) AddChat(chat *tgbotapi.Chat) {
	switch {
	case chat.IsPrivate():
		c.private = append(c.private, chat.ID)
	case chat.IsGroup() || chat.IsSuperGroup():
		c.group = append(c.group, chat.ID)
	}
}
