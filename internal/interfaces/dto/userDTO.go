package dto

type UserDTO struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	StudentID *int64 `json:"student_id,omitempty"`
	TrainerID *int64 `json:"trainer_id,omitempty"`
}
