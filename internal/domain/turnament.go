package domain

import "time"

type Tournament struct {
	ID       int64     `gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"type:text;not null"`
	Room     string    `gorm:"type:text;not null"`
	Datetime time.Time `gorm:"not null"`
	TeamID   int       `gorm:"not null"`
	Team     Team      `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
