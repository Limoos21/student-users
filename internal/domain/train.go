package domain

import "time"

type Train struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Type      string    `gorm:"type:text;not null"`
	Room      string    `gorm:"type:text;not null"`
	Datetime  time.Time `gorm:"not null"`
	TrainerID int64     `gorm:"not null"`
	TeamID    int       `gorm:"not null"`
	Trainer   Trainer   `gorm:"foreignKey:TrainerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Team      Team      `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
