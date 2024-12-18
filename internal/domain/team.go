package domain

type Team struct {
	ID     int    `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"type:varchar(255);not null"`
	League string `gorm:"type:varchar(255);not null"`
}
