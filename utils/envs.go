package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Envs struct {
	BotToken string
	DbUri    string
}

func GetEnvs() (Envs, error) {
	godotenv.Load()
	var envs Envs

	envs.BotToken = os.Getenv("BOT_TOKEN")
	envs.DbUri = os.Getenv("DB_URI")

	return envs, nil
}
