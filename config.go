package main

import (
	"github.com/pkg/errors"
	"math/rand"
	"os"
)

type configuration struct {
	logLevel configParameter

	serverHost configParameter
	serverPort configParameter

	useWebhookSecureKey configParameter
	webhookSecureKey    configParameter
}

func (c *configuration) loadFromEnv() error {

	if err := setFromEnv(&conf.logLevel, "LOG_LEVEL",
		"debug", false, true); err != nil {
		return err
	}

	if err := setFromEnv(&conf.serverHost, "SERVER_HOST",
		"0.0.0.0", false, true); err != nil {
		return err
	}

	if err := setFromEnv(&conf.serverPort, "SERVER_PORT",
		"80", false, true); err != nil {
		return err
	}

	if err := setFromEnv(&conf.useWebhookSecureKey, "USE_SECURE_KEY",
		"", true, false); err != nil {
		return err
	}

	if conf.useWebhookSecureKey.Bool() {
		if err := setFromEnv(&conf.webhookSecureKey, "SECURE_KEY",
			"", true, false); err != nil {

			keyFileName := "sec.key"
			var key string
			letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
			b := make([]rune, 64)
			for i := range b {
				b[i] = letters[rand.Intn(len(letters))]
			}
			key = string(b)

			if _, err := os.Stat(keyFileName); os.IsNotExist(err) {
				if err := os.WriteFile(keyFileName, []byte(string(b)), 0777); err != nil {
					return errors.WithMessage(err, "failed to write security key to file")
				}
			} else {
				b, err := os.ReadFile(keyFileName)
				if err == nil {
					key = string(b)
				}
			}

			conf.webhookSecureKey = configParameter(key)
			log.Debugf("Service secure key: %s", key)
		}
	}

	return nil
}

func (c *configuration) serverAddress() string {
	return c.serverHost.String() + ":" + c.serverPort.String()
}

var conf = &configuration{}
