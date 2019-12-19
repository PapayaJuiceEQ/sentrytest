package main

import (
	"errors"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
)

const (
	sentryDSN = "http://90eadccc602747e7ad90e4bd65751ae9@35.245.29.149:9000/2"
)

func init() {
	hook, err := logrus_sentry.NewSentryHook(sentryDSN, []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
	})
	if err != nil {
		log.Fatalf("error setting up sentry hook: %v", err)
	}

	log.AddHook(hook)
}

func main() {
	log.WithField("WhatsWrong", "idk dude").Error(errors.New("this is a bad error"))
}
