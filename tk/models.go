package tk

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Username   string `gorm:"unique_index"`
	Datapoints []Datapoint
}

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

func performMigrations() {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Datapoint{})
	db.AutoMigrate(&SkillLevel{})
}
