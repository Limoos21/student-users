package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"stud-trener/internal/application"
	"stud-trener/internal/interfaces/dto"
)

type TrainerHandler struct {
	useCase application.TrainerUseCaseInterface
	logger  *zap.Logger
}

func NewTrainerHandler(r *gin.RouterGroup, useCase application.TrainerUseCaseInterface, logger *zap.Logger) {
	handler := &TrainerHandler{useCase, logger}
	r.POST("/create", handler.CreateTrainer)
	r.POST("/update/:id", handler.UpdateTrainer)
	r.POST("/delete/:id", handler.DeleteTrainer)
	r.GET("/", handler.GetTrainers)
	r.GET("/:id", handler.GetTrainer)
}

func (handler *TrainerHandler) CreateTrainer(c *gin.Context) {
	var data dto.TrainerDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTrainer, err := handler.useCase.CreateTrainer(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTrainer)
}

func (handler *TrainerHandler) UpdateTrainer(c *gin.Context) {
	var data dto.TrainerDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	updatedTrainer, err := handler.useCase.UpdateTrainer(int64(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTrainer)
}

func (handler *TrainerHandler) DeleteTrainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = handler.useCase.DeleteTrainer(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trainer deleted"})
}

func (handler *TrainerHandler) GetTrainers(c *gin.Context) {
	trainers, err := handler.useCase.GetAllTrainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trainers)
}

func (handler *TrainerHandler) GetTrainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	trainer, err := handler.useCase.GetTrainerByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trainer)
}
