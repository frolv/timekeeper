package tk

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"timekeeper/lib/osrs"
	"timekeeper/lib/tkerr"
)

// Create a new datapoint for the account with the specified username.
func UpdateAccount(username string) (*Account, error) {
	name, err := canonicalizeUsername(username)
	if err != nil {
		return nil, err
	}

	var account Account
	first := false

	if err := db.First(&account, "username = ?", name).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			account = Account{
				Username:    name,
				DisplayName: username,
			}
			first = true
		} else {
			fmt.Println(err.Error())
			return nil, err
		}
	} else {
		var dp Datapoint
		db.Where("account_id = ?", account.ID).Last(&dp)

		if time.Since(dp.CreatedAt) < time.Duration(30*time.Second) {
			return nil, tkerr.Create(tkerr.RecentUpdate)
		}
	}

	if err := createDatapoint(&account, first); err != nil {
		return nil, err
	}

	return &account, nil
}

// Fetch current account stats from the OSRS API and create
// a new datapoint for the specified account.
func createDatapoint(account *Account, first bool) error {
	skills, err := osrs.HiscoreLookup(account.Username)
	if err != nil {
		return err
	}

	if first {
		db.Create(account)
	}

	dp := Datapoint{AccountID: account.ID}
	db.Create(&dp)

	for _, sk := range skills {
		sl := SkillLevel{
			DatapointID: dp.ID,
			SkillID:     sk.ID,
			Experience:  sk.Experience,
			Level:       sk.Level,
			Rank:        sk.Rank,
			Hours:       sk.Hours,
		}
		db.Create(&sl)
	}

	return nil
}
