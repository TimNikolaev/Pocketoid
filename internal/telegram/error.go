package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errFailToSave   = errors.New("fail to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, replyUnknownError)

	switch err {
	case errInvalidURL:
		msg.Text = replyInvalidURL
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = replyUnauthorized
		b.bot.Send(msg)
	case errFailToSave:
		msg.Text = replyFailToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
