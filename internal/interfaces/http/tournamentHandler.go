package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"stud-trener/internal/application"
	"stud-trener/internal/interfaces/dto"
)

type TournamentHandler struct {
	useCase application.TournamentUseCaseInterface
	logger  *zap.Logger
}

func NewTournamentHandler(r *gin.RouterGroup, useCase application.TournamentUseCaseInterface, logger *zap.Logger) {
	handler := &TournamentHandler{useCase, logger}
	r.POST("/create", handler.CreateTournament)
	r.POST("/update/:id", handler.UpdateTournament)
	r.POST("/delete/:id", handler.DeleteTournament)
	r.GET("/", handler.GetTournaments)
	r.GET("/:id", handler.GetTournament)
}

func (handler *TournamentHandler) CreateTournament(c *gin.Context) {
	var data dto.TournamentDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTournament, err := handler.useCase.CreateTournament(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTournament)
}

func (handler *TournamentHandler) UpdateTournament(c *gin.Context) {
	var data dto.TournamentDTO
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
	updatedTournament, err := handler.useCase.UpdateTournament(int64(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTournament)
}

func (handler *TournamentHandler) DeleteTournament(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	err = handler.useCase.DeleteTournament(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tournament deleted"})
}

func (handler *TournamentHandler) GetTournaments(c *gin.Context) {
	tournaments, err := handler.useCase.GetAllTournaments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tournaments)
}

func (handler *TournamentHandler) GetTournament(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	tournament, err := handler.useCase.GetTournamentByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tournament)
}
