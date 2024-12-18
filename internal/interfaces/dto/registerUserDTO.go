package dto

type RegisterUserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	StudentID *int64 `json:"student_id,omitempty"`
	TrainerID *int64 `json:"trainer_id,omitempty"`
}
