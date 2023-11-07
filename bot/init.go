package bot

import (
	"taksa/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot tgbotapi.BotAPI

func Init() <-chan error {
	c := make(chan error)

	envs, err := utils.GetEnvs()

	if err != nil {
		c <- err
	}

	go func() {
		newBot, err := tgbotapi.NewBotAPI(envs.BotToken)

		if !bot.Self.IsBot {
			c <- err
		}

		bot = *newBot

		c <- nil
	}()

	return c
}

func GetBot() *tgbotapi.BotAPI {
	return &bot
}
