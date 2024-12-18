package dto

type StudentDTO struct {
	ID     *int   `json:"id,omitempty"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	TeamID int    `json:"team_id"`
}
