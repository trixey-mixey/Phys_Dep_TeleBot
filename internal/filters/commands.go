package filters

import "github.com/go-telegram/bot/models"

const (
	start = "/start"
	count = "/count"
)

func IsStart(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == start
}

func IsCount(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == count
}
