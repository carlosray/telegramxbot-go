package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"regexp"
	"strings"
	"telegramxbot/internal/util"
	"time"
)

type HueHandler struct {
	minMessages int
	maxMessages int
	msgCounts   struct {
		actual    map[int64]int
		expecting map[int64]int
	}
}

var vowels = map[rune]rune{
	'а': 'я',
	'у': 'ю',
	'о': 'е',
	'ы': 'и',
	'и': 'и',
	'э': 'е',
	'я': 'я',
	'ю': 'ю',
	'е': 'е',
	'ё': 'ё',
}

func (h *HueHandler) Setup(props map[string]interface{}) {
	h.minMessages = props["min"].(int)
	h.maxMessages = props["max"].(int)
	h.msgCounts.actual = make(map[int64]int)
	h.msgCounts.expecting = make(map[int64]int)
}

func (h *HueHandler) Handle(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (bool, error) {
	var msg *tgbotapi.Message
	var ok bool
	if msg, ok = util.GetNonCommandMessage(update); !ok || !h.needReply(msg.Chat.ID) {
		return false, nil
	}

	log.Infof("Processing 'hue' from chat %d", msg.Chat.ID)

	message := getHueMessage(msg.Text)

	if len(message) <= 0 {
		return false, nil
	}

	msgToSend := tgbotapi.NewMessage(msg.Chat.ID, message)
	msgToSend.ReplyToMessageID = msg.MessageID

	if _, err := bot.Send(msgToSend); err != nil {
		return false, err
	}
	return true, nil
}

func (h *HueHandler) needReply(chat int64) bool {
	h.msgCounts.actual[chat] += 1
	a := h.msgCounts.actual[chat]
	var e int
	var ok bool
	if e, ok = h.msgCounts.expecting[chat]; !ok {
		e = getRandom(h.minMessages, h.maxMessages)
		h.msgCounts.expecting[chat] = e
	}
	if a >= e {
		h.msgCounts.expecting[chat] = getRandom(h.minMessages, h.maxMessages)
		h.msgCounts.actual[chat] = 0
		return true
	}
	return false
}

func getRandom(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func getHueMessage(input string) string {
	//get last row
	rows := strings.Split(input, "\n")
	s := rows[len(rows)-1]
	//get last word
	re := regexp.MustCompile("\\s*(\\s|,|!|\\.|\\)|\\(|\\?)\\s*")
	words := re.Split(s, -1)
	s = words[len(words)-1]
	//check if cyrillic
	re = regexp.MustCompile("[а-яёА-ЯЁ]+")
	if re.MatchString(s) {
		return concatHue(strings.ToLower(s))
	}
	return ""
}

func concatHue(s string) (r string) {
	var replaced = false
	for _, c := range s {
		if replaced {
			r += string(c)
		}
		if v, ok := vowels[c]; ok && !replaced {
			r += string(v)
			replaced = true
		}
	}
	if !replaced {
		return ""
	}
	return "ху" + r
}
