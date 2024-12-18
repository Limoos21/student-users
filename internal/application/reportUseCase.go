package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type ReportUseCase struct {
	Repository repository.ReportRepositoryInterface
	logger     *zap.Logger
}

type ReportUseCaseInterface interface {
	GetAttendanceReport(studentID int64) (dto.AttendanceReportDTO, error)
	GetCompetitionReport(studentID int64) ([]dto.CompetitionReportDTO, error)
	GetTrainingScheduleReport(studentID int64) ([]dto.TrainingScheduleReportDTO, error)
}

func NewReportUseCase(repository repository.ReportRepositoryInterface, logger *zap.Logger) *ReportUseCase {
	return &ReportUseCase{Repository: repository, logger: logger}
}

// Получение отчета по посещаемости
func (r *ReportUseCase) GetAttendanceReport(studentID int64) (dto.AttendanceReportDTO, error) {
	report, err := r.Repository.GetAttendanceReport(studentID)
	if err != nil {
		r.logger.Error("Error fetching attendance report", zap.Int64("student_id", studentID), zap.Error(err))
		return dto.AttendanceReportDTO{}, err
	}

	// Преобразуем доменную модель в DTO
	return dto.AttendanceReportDTO{
		TrainerName:       report.TrainerName,
		StudentName:       report.StudentName,
		TeamName:          report.TeamName,
		AttendedTrainings: report.AttendedTrainings,
		MissedTrainings:   report.MissedTrainings,
	}, nil
}

// Получение отчета по участию в соревнованиях
func (r *ReportUseCase) GetCompetitionReport(studentID int64) ([]dto.CompetitionReportDTO, error) {
	reports, err := r.Repository.GetCompetitionReport(studentID)
	if err != nil {
		r.logger.Error("Error fetching competition report", zap.Int64("student_id", studentID), zap.Error(err))
		return nil, err
	}

	// Преобразуем список доменных моделей в DTO
	competitionReports := make([]dto.CompetitionReportDTO, len(reports))
	for i, report := range reports {
		competitionReports[i] = dto.CompetitionReportDTO{
			TrainerName:      report.TrainerName,
			StudentName:      report.StudentName,
			TeamName:         report.TeamName,
			CompetitionName:  report.CompetitionName,
			CompetitionDate:  report.CompetitionDate,
			CompetitionPlace: report.CompetitionPlace,
		}
	}
	return competitionReports, nil
}

// Получение отчета по расписанию тренировок
func (r *ReportUseCase) GetTrainingScheduleReport(studentID int64) ([]dto.TrainingScheduleReportDTO, error) {
	reports, err := r.Repository.GetTrainingScheduleReport(studentID)
	if err != nil {
		r.logger.Error("Error fetching training schedule report", zap.Int64("student_id", studentID), zap.Error(err))
		return nil, err
	}

	// Преобразуем список доменных моделей в DTO
	trainingScheduleReports := make([]dto.TrainingScheduleReportDTO, len(reports))
	for i, report := range reports {
		trainingScheduleReports[i] = dto.TrainingScheduleReportDTO{
			TrainerName:  report.TrainerName,
			StudentName:  report.StudentName,
			TeamName:     report.TeamName,
			TrainingDate: report.TrainingDate,
			Room:         report.Room,
			StartTime:    report.StartTime,
		}
	}
	return trainingScheduleReports, nil
}
