package main

import (
	"os"

	"github.com/altfoxie/labrys-bot/internal/bot"
	"github.com/altfoxie/labrys-bot/internal/storage"

	"github.com/sirupsen/logrus"
)

func main() {
	s, err := storage.New("storage.db")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize storage.")
	}

	b, err := bot.New(s, os.Getenv("TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize bot.")
	}

	if err = b.Start(); err != nil {
		logrus.WithError(err).Fatal("Failed to start bot.")
	}
}
