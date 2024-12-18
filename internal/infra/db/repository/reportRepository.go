package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

// Репозиторий для отчетов
type ReportRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewReportRepository(db *gorm.DB, logger *zap.Logger) *ReportRepository {
	return &ReportRepository{db: db, logger: logger}
}

// Интерфейс репозитория отчетов
type ReportRepositoryInterface interface {
	GetAttendanceReport(studentID int64) (*domain.AttendanceReport, error)
	GetCompetitionReport(studentID int64) ([]domain.CompetitionReport, error)
	GetTrainingScheduleReport(studentID int64) ([]domain.TrainingScheduleReport, error)
}

// Получение отчета по посещаемости для ученика
func (r *ReportRepository) GetAttendanceReport(studentID int64) (*domain.AttendanceReport, error) {
	var report domain.AttendanceReport
	r.logger.Info("Fetching attendance report for student", zap.Int64("student_id", studentID))
	query := `
		SELECT 
    tr.name AS "TrainerName",
    st.name AS "StudentName",
    tm.name AS "TeamName",
    COUNT(trn.id) AS "AttendedTrainings",
    COUNT(all_trains.id) - COUNT(trn.id) AS "MissedTrainings"
FROM work.student st
JOIN work.team tm ON st.team_id = tm.id
JOIN work.team_trainer tt ON tt.id_team = tm.id
JOIN work.trainer tr ON tr.id = tt.id_trainer
LEFT JOIN work.train all_trains ON all_trains.id_team = tm.id AND all_trains.id_trainer = tr.id AND all_trains.datetime < CURRENT_TIMESTAMP
LEFT JOIN work.train trn ON trn.id_trainer = tr.id AND trn.id_team = tm.id AND trn.datetime < CURRENT_TIMESTAMP
WHERE st.id = ?
GROUP BY tr.name, st.name, tm.name;`
	result := r.db.Raw(query, studentID).Scan(&report)
	if result.Error != nil {
		r.logger.Error("Error fetching attendance report", zap.Int64("student_id", studentID), zap.Error(result.Error))
		return nil, result.Error
	}
	return &report, nil
}

// Получение отчета по участию в соревнованиях
func (r *ReportRepository) GetCompetitionReport(studentID int64) ([]domain.CompetitionReport, error) {
	var reports []domain.CompetitionReport
	r.logger.Info("Fetching competition report for student", zap.Int64("student_id", studentID))
	query := `
		SELECT 
			tr.name AS "TrainerName",
			st.name AS "StudentName",
			tm.name AS "TeamName",
			t.name AS "CompetitionName",
			t.datetime AS "CompetitionDate",
			t.room AS "CompetitionPlace"
		FROM work.student st
		JOIN work.team tm ON st.team_id = tm.id
		JOIN work.team_trainer tt ON tt.id_team = tm.id
		JOIN work.trainer tr ON tr.id = tt.id_trainer
		JOIN work.tournament t ON t.id_team = tm.id
		WHERE st.id = ?
		ORDER BY t.datetime`
	result := r.db.Raw(query, studentID).Scan(&reports)
	if result.Error != nil {
		r.logger.Error("Error fetching competition report", zap.Int64("student_id", studentID), zap.Error(result.Error))
		return nil, result.Error
	}
	return reports, nil
}

// Получение отчета по расписанию тренировок
func (r *ReportRepository) GetTrainingScheduleReport(studentID int64) ([]domain.TrainingScheduleReport, error) {
	var reports []domain.TrainingScheduleReport
	r.logger.Info("Fetching training schedule report for student", zap.Int64("student_id", studentID))
	query := `
		SELECT 
			tr.name AS "TrainerName",
			st.name AS "StudentName",
			tm.name AS "TeamName",
			trn.datetime AS "TrainingDate",
			trn.room AS "Room",
			trn.datetime AS "StartTime"
		FROM work.student st
		JOIN work.team tm ON st.team_id = tm.id
		JOIN work.team_trainer tt ON tt.id_team = tm.id
		JOIN work.trainer tr ON tr.id = tt.id_trainer
		JOIN work.train trn ON trn.id_trainer = tr.id AND trn.id_team = tm.id
		WHERE st.id = ?
		ORDER BY trn.datetime`
	result := r.db.Raw(query, studentID).Scan(&reports)
	if result.Error != nil {
		r.logger.Error("Error fetching training schedule report", zap.Int64("student_id", studentID), zap.Error(result.Error))
		return nil, result.Error
	}
	return reports, nil
}
