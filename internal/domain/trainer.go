package domain

type Trainer struct {
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:text;not null"`
	Age  int    `gorm:"not null"`
}
