package dto

type TeamDTO struct {
	ID     *int   `json:"id,omitempty"`
	Name   string `json:"name"`
	League string `json:"league"`
}
