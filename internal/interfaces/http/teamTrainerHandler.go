package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"stud-trener/internal/application"
	"stud-trener/internal/interfaces/dto"
)

type TeamTrainerHandler struct {
	useCase application.TeamTrainerUseCaseInterface
	logger  *zap.Logger
}

func NewTeamTrainerHandler(r *gin.RouterGroup, useCase application.TeamTrainerUseCaseInterface, logger *zap.Logger) {
	handler := &TeamTrainerHandler{useCase, logger}
	r.POST("/create", handler.CreateTeamTrainer)
	r.POST("/update/:id", handler.UpdateTeamTrainer)
	r.POST("/delete/:id", handler.DeleteTeamTrainer)
	r.GET("/", handler.GetTeamTrainers)
	r.GET("/:id", handler.GetTeamTrainer)
}

func (handler *TeamTrainerHandler) CreateTeamTrainer(c *gin.Context) {
	var data dto.TeamTrainerDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTeamTrainer, err := handler.useCase.CreateTeamTrainer(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTeamTrainer)
}

func (handler *TeamTrainerHandler) UpdateTeamTrainer(c *gin.Context) {
	var data dto.TeamTrainerDTO
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
	updatedTeamTrainer, err := handler.useCase.UpdateTeamTrainer(int64(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTeamTrainer)
}

func (handler *TeamTrainerHandler) DeleteTeamTrainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = handler.useCase.DeleteTeamTrainer(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team-Trainer relationship deleted"})
}

func (handler *TeamTrainerHandler) GetTeamTrainers(c *gin.Context) {
	teamTrainers, err := handler.useCase.GetAllTeamTrainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teamTrainers)
}

func (handler *TeamTrainerHandler) GetTeamTrainer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	teamTrainer, err := handler.useCase.GetTeamTrainerByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teamTrainer)
}
