package tk

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Username    string `gorm:"unique_index"`
	DisplayName string `gorm:"size:16"`
	Datapoints  []Datapoint
}

type Datapoint struct {
	ID          uint      `gorm:"primary_key"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	AccountID   uint       `gorm:"type:bigint REFERENCES accounts(id)"`
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
