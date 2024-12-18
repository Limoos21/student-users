package domain

type Student struct {
	ID     int    `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"type:varchar(255);not null"`
	Age    int    `gorm:"not null"`
	Height int    `gorm:"not null"`
	Weight int    `gorm:"not null"`
	TeamID int    `gorm:"not null"`
	Team   Team   `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
