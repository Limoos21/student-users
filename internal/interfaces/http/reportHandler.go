package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"stud-trener/internal/application"
)

type ReportHandler struct {
	useCase application.ReportUseCaseInterface
	logger  *zap.Logger
}

func NewReportHandler(r *gin.RouterGroup, useCase application.ReportUseCaseInterface, logger *zap.Logger) {
	handler := &ReportHandler{useCase, logger}
	r.GET("/attendance", handler.GetAttendanceReport)
	r.GET("/competition", handler.GetCompetitionReport)
	r.GET("/schedule", handler.GetTrainingSchedule)
}

// Получение отчета по посещаемости
func (handler *ReportHandler) GetAttendanceReport(c *gin.Context) {
	studentIDParam := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student_id parameter"})
		return
	}

	report, err := handler.useCase.GetAttendanceReport(int64(studentID))
	if err != nil {
		handler.logger.Error("Error fetching attendance report", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

// Получение отчета по соревнованиям
func (handler *ReportHandler) GetCompetitionReport(c *gin.Context) {
	studentIDParam := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student_id parameter"})
		return
	}

	report, err := handler.useCase.GetCompetitionReport(int64(studentID))
	if err != nil {
		handler.logger.Error("Error fetching competition report", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

// Получение расписания тренировок
func (handler *ReportHandler) GetTrainingSchedule(c *gin.Context) {
	studentIDParam := c.Query("student_id")
	studentID, err := strconv.Atoi(studentIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student_id parameter"})
		return
	}

	report, err := handler.useCase.GetTrainingScheduleReport(int64(studentID))
	if err != nil {
		handler.logger.Error("Error fetching training schedule", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
