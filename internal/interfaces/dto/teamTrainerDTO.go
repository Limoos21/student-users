package dto

type TeamTrainerDTO struct {
	ID        *int64 `json:"id,omitempty"`
	TeamID    int    `json:"team_id"`
	TrainerID int64  `json:"trainer_id"`
}
