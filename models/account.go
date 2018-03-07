package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"timekeeper/lib/osrs"
	"unicode"
)

type Account struct {
	gorm.Model
	Username   string `gorm:"unique_index"`
	Datapoints []Datapoint
}

type UpdateError struct {
	Type    int
	Message string
}

const (
	UEInvalidAccount  = osrs.HEInvalidAccount
	UEAPIError        = osrs.HEAPIError
	UEInvalidUsername = iota
	UERecentUpdate
)

func (e *UpdateError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Type, e.Message)
}

// Create a new datapoint for the account with the specified username.
func UpdateAccount(username string) error {
	if !validUsername(username) {
		return &UpdateError{UEInvalidUsername, "Invalid username"}
	}

	var account Account
	first := false

	if err := db.First(&account, "username = ?", username).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			account = Account{Username: username}
			first = true
		} else {
			fmt.Println(err.Error())
			return err
		}
	} else {
		var dp Datapoint
		db.Where("account_id = ?", account.ID).Last(&dp)

		if time.Since(dp.CreatedAt) < time.Duration(30*time.Second) {
			return &UpdateError{
				UERecentUpdate,
				"Account updated less than 30s ago",
			}
		}
	}

	if err := createDatapoint(&account, first); err != nil {
		if e, ok := err.(*osrs.HiscoreError); ok {
			return &UpdateError{e.Type, e.Message}
		} else {
			fmt.Println(err.Error())
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