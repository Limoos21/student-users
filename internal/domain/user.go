package domain

type User struct {
	ID           int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string `gorm:"unique;not null" json:"username"`
	PasswordHash string `gorm:"not null" json:"password_hash"`
	Role         string `gorm:"not null;check:role IN ('student', 'trainer')" json:"role"`
	StudentID    *int64 `gorm:"index" json:"student_id,omitempty"`
	TrainerID    *int64 `gorm:"index" json:"trainer_id,omitempty"`
}
