package domain

type AttendanceReport struct {
	TrainerName       string `json:"trainer_name"`
	StudentName       string `json:"student_name"`
	TeamName          string `json:"team_name"`
	AttendedTrainings int    `json:"attended_trainings"`
	MissedTrainings   int    `json:"missed_trainings"`
}
