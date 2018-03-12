package tk

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"timekeeper/lib/tkerr"
)

// Look up an account by username.
func GetAccount(username string) (*Account, error) {
	var account Account

	if err := db.First(&account, "username = ?", username).Error; err != nil {
		return nil, tkerr.Create(tkerr.UntrackedAccount)
	} else {
		return &account, nil
	}
}

// Fetches the latest datapoint for an account.
func LatestDatapoint(acc *Account) (*Datapoint, error) {
	var dp Datapoint

	if acc == nil {
		return nil, errors.New("No account provided")
	}

	db.Preload("SkillLevels", func(db *gorm.DB) *gorm.DB {
		return db.Order(`"skill_levels"."id" ASC`)
	}).Where("account_id = ?", acc.ID).Last(&dp)

	return &dp, nil
}

// Fetches the first and last datapoints within a specified time period.
func BoundaryPoints(acc *Account, start time.Time, end time.Time) ([]Datapoint, error) {
	if acc == nil {
		return nil, errors.New("No account provided")
	}

	dps := make([]Datapoint, 2)

	// Fetch the first datapoint in the period with its skills, and then the
	// most recent datapoint. This takes 4 queries, which isn't ideal, but nor
	// is it bad enough that it needs to be optimized (at least for now).
	// We've got redis in front of this anyway.
	err := db.Preload("SkillLevels").Where(
		"account_id = ? AND created_at >= ? AND created_at <= ?",
		acc.ID, start, end,
	).First(&dps[0]).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, tkerr.Create(tkerr.NoPointsInPeriod)
		} else {
			fmt.Println(err)
			return nil, err
		}
	}

	db.Preload("SkillLevels", func(db *gorm.DB) *gorm.DB {
		return db.Order(`"skill_levels"."id" ASC`)
	}).Where("account_id = ? AND created_at <= ?", acc.ID, end).Last(&dps[1])

	return dps, nil
}
