package db

import "taksa/utils"

var CreatingAccounts []int64

func AddNewAccountCreating(userId int64) {
	CreatingAccounts = append(CreatingAccounts, userId)
}

func CheckExistAccountCreating(userId int64) bool {
	for i := range CreatingAccounts {
		if CreatingAccounts[i] == userId {
			return true
		}
	}
	return false
}

func RemoveAccountCreating(userId int64) {
	for i := range CreatingAccounts {
		if CreatingAccounts[i] == userId {
			updatedAccounts := utils.DeleteSliceElement(CreatingAccounts, i)
			CreatingAccounts = updatedAccounts
		}
	}
}
