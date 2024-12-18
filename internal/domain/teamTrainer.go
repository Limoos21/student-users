package domain

type TeamTrainer struct {
	ID        int64   `gorm:"primaryKey;autoIncrement"`
	TeamID    int     `gorm:"not null"`
	TrainerID int64   `gorm:"not null"`
	Team      Team    `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Trainer   Trainer `gorm:"foreignKey:TrainerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
