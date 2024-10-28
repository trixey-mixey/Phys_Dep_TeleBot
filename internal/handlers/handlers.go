package handlers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/trixey-mixey/Phys_Dep_TeleBot/algho"
)

var userState = make(map[int64]string)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text
	state, exists := userState[chatID]

	if exists && state == "waiting_for_numbers" {
		convertStringAndSendTable(text, ctx, b, update)
		delete(userState, chatID)
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "Неизвестная команда",
		})
	}

}

func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Добро пожаловать в Unn Hack бот. Выберите опцию",
	})
}

func Count(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите измерения через пробел, разделяя дробную часть точкой.\nНапример: 12.2 12.4 12.3",
	})

	userState[update.Message.Chat.ID] = "waiting_for_numbers"

}

func convertStringAndSendTable(str string, ctx context.Context, b *bot.Bot, update *models.Update) error {
	if !strings.Contains(str, " ") {
		err := errors.New("разделите измерения пробелом")
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return err
	}
	strSlice := strings.Split(str, " ")
	var floatSlice []float64
	for _, s := range strSlice {
		floatEl, _ := strconv.ParseFloat(s, 64)
		floatSlice = append(floatSlice, floatEl)
	}
	avg, err := algho.GetAverage(floatSlice...)

	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return err
	}

	avgStr := strconv.FormatFloat(avg, 'f', -1, 64)

	avgMinusEl, _ := algho.GetAverageMinusEl(floatSlice...)
	var avgMinusElStrSlice []string
	for _, num := range avgMinusEl {
		avgMinusElStrSlice = append(avgMinusElStrSlice, strconv.FormatFloat(num, 'f', -1, 64))
	}

	sqr, _ := algho.GetSquare(floatSlice...)
	var sqrSlice []string
	for _, num := range sqr {
		sqrSlice = append(sqrSlice, strconv.FormatFloat(num, 'f', -1, 64))
	}

	SO, _ := algho.GetSO(floatSlice...)
	SOString := strconv.FormatFloat(SO, 'f', -1, 64)

	instrErr := algho.GetInstrErr(floatSlice...)
	instrErrString := strconv.FormatFloat(instrErr, 'f', -1, 64)

	randErr, _ := algho.GetRandErr(floatSlice...)
	randErrString := strconv.FormatFloat(randErr, 'f', -1, 64)

	fullErr, _ := algho.GetFullErr(floatSlice...)
	fullErrString := strconv.FormatFloat(fullErr, 'f', -1, 64)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("d ср: %s\n------------------------\nd`-d: %s\n------------------------\n(d`-d)^2: %s\n------------------------\nSo: %s\n------------------------\n△d,сл: %s\n------------------------\n△d,пр: %s\n------------------------\n△d: %s ", avgStr, strings.Join(avgMinusElStrSlice, ", "), strings.Join(sqrSlice, ", "), SOString, randErrString, instrErrString, fullErrString),
	})

	return nil
}
