package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"stud-trener/internal/application"
	"stud-trener/internal/interfaces/dto"
)

type TrainHandler struct {
	useCase application.TrainUseCaseInterface
	logger  *zap.Logger
}

func NewTrainHandler(r *gin.RouterGroup, useCase application.TrainUseCaseInterface, logger *zap.Logger) {
	handler := &TrainHandler{useCase, logger}
	r.POST("/create", handler.CreateTrain)
	r.POST("/update/:id", handler.UpdateTrain)
	r.POST("/delete/:id", handler.DeleteTrain)
	r.GET("/", handler.GetTrains)
	r.GET("/:id", handler.GetTrain)
}

func (handler *TrainHandler) CreateTrain(c *gin.Context) {
	var data dto.TrainDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTrain, err := handler.useCase.CreateTrain(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTrain)
}

func (handler *TrainHandler) UpdateTrain(c *gin.Context) {
	var data dto.TrainDTO
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
	updatedTrain, err := handler.useCase.UpdateTrain(int64(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTrain)
}

func (handler *TrainHandler) DeleteTrain(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = handler.useCase.DeleteTrain(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Training session deleted"})
}

func (handler *TrainHandler) GetTrains(c *gin.Context) {
	trains, err := handler.useCase.GetAllTrains()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trains)
}

func (handler *TrainHandler) GetTrain(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	train, err := handler.useCase.GetTrainByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, train)
}
