package domain

import "time"

type TrainingScheduleReport struct {
	TrainerName  string    `json:"trainer_name"`
	StudentName  string    `json:"student_name"`
	TeamName     string    `json:"team_name"`
	TrainingDate time.Time `json:"training_date"`
	Room         string    `json:"room"`
	StartTime    time.Time `json:"start_time"`
}
