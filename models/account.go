package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"unicode"
)

type Account struct {
	gorm.Model
	Username string `gorm:"unique_index"`
}

type AccountError struct {
	Type    int
	Message string
}

const (
	AEInvalidUser = iota
)

func (e *AccountError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Type, e.Message)
}

func UpdateAccount(username string) error {
	if !validUsername(username) {
		return &AccountError{AEInvalidUser, "Invalid username"}
	}

	var account Account
	err := db.First(&account, "username = ?", username).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			account = Account{Username: username}
			db.Create(&account)
		} else {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func validUsername(username string) bool {
	valid := true
	for _, c := range username {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) && c != '_' {
			valid = false
			break
		}
	}

	return valid && len(username) > 0 && len(username) <= 12
}
