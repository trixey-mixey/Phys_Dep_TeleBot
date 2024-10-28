package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/trixey-mixey/Phys_Dep_TeleBot/internal/filters"
	"github.com/trixey-mixey/Phys_Dep_TeleBot/internal/handlers"
	"github.com/trixey-mixey/Phys_Dep_TeleBot/pkg/systems"
)

func main() {
	BOT_TOKEN := systems.BotToken()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.DefaultHandler),
	}
	b, err := bot.New(BOT_TOKEN, opts...)
	b.RegisterHandlerMatchFunc(filters.IsStart, handlers.Start)
	b.RegisterHandlerMatchFunc(filters.IsCount, handlers.Count)

	if nil != err {
		panic(err)
	}
	b.Start(ctx)
}
