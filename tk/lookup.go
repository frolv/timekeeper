package tk

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Look up an account by username.
func GetAccount(username string) (*Account, error) {
	var account Account

	if err := db.First(&account, "username = ?", username).Error; err != nil {
		return nil, errors.New("Account has not been tracked")
	} else {
		return &account, nil
	}
}

// Fetches the latest datapoint for an account.
func GetLatestDatapoint(acc *Account) (*Datapoint, error) {
	var dp Datapoint

	if acc == nil {
		return nil, errors.New("No account provided")
	}

	db.Preload("SkillLevels", func(db *gorm.DB) *gorm.DB {
		return db.Order(`"skill_levels"."id" ASC`)
	}).Where("account_id = ?", acc.ID).Last(&dp)

	return &dp, nil
}
