package dto

import "time"

type TournamentDTO struct {
	ID       *int64    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Room     string    `json:"room"`
	Datetime time.Time `json:"datetime"`
	TeamID   int       `json:"team_id"`
}
