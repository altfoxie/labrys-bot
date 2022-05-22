package main

import (
	"os"
	"strconv"

	"github.com/altfoxie/labrys-bot/bot"
	"github.com/altfoxie/labrys-bot/storage"

	"github.com/sirupsen/logrus"
)

func main() {
	s, err := storage.New("storage.db")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize storage.")
	}

	id, _ := strconv.ParseInt(os.Getenv("CHANNEL"), 10, 64)
	if id == 0 {
		logrus.Warn("Channel ID is not set.")
	}

	b, err := bot.New(bot.Params{
		Token:     os.Getenv("TOKEN"),
		Storage:   s,
		ChannelID: id,
	})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize bot.")
	}

	if err = b.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start bot.")
	}
}
