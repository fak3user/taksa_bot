package bot

import (
	"taksa/db"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartPrivateChat(user *tgbotapi.User) (bool, error) {
	newOrExistingUser, isNew, err := db.AddUserAndCheckIfExist(user)
	if err != nil {
		return false, err
	}

	spew.Dump(newOrExistingUser)

	return isNew, nil
}
