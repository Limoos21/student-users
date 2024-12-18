package dto

import "time"

type CompetitionReportDTO struct {
	TrainerName      string    `json:"trainer_name"`
	StudentName      string    `json:"student_name"`
	TeamName         string    `json:"team_name"`
	CompetitionName  string    `json:"competition_name"`
	CompetitionDate  time.Time `json:"competition_date"`
	CompetitionPlace string    `json:"competition_place"`
}
