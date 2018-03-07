package models

import (
	"github.com/jinzhu/gorm"
	"timekeeper/lib/osrs"
)

type Datapoint struct {
	gorm.Model
	AccountID   uint `gorm:"type:bigint REFERENCES accounts(id)"`
	Account     Account
	SkillLevels []SkillLevel
}

type SkillLevel struct {
	gorm.Model
	DatapointID uint `gorm:"type:bigint REFERENCES datapoints(id)"`
	Datapoint   Datapoint
	SkillID     uint
	Experience  int `gorm:"type:bigint"`
	Level       int
	Rank        int
	Hours       float64
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
