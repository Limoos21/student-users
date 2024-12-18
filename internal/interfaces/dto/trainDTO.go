package dto

import "time"

type TrainDTO struct {
	ID        *int64    `json:"id,omitempty"`
	Type      string    `json:"type"`
	Room      string    `json:"room"`
	Datetime  time.Time `json:"datetime"`
	TrainerID int64     `json:"trainer_id"`
	TeamID    int       `json:"team_id"`
}
