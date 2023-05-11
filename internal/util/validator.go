package util

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func GetCommandMessage(update *tgbotapi.Update) (*tgbotapi.Message, bool) {
	if message, ok := GetMessage(update); ok && message.IsCommand() {
		return message, true
	}
	return nil, false
}

func GetNonCommandMessage(update *tgbotapi.Update) (*tgbotapi.Message, bool) {
	if message, ok := GetMessage(update); ok && !message.IsCommand() {
		return message, true
	}
	return nil, false
}

func GetMessage(update *tgbotapi.Update) (*tgbotapi.Message, bool) {
	var msg *tgbotapi.Message
	if msg = update.Message; msg == nil || msg.Chat == nil {
		return nil, false
	}
	return msg, true
}
