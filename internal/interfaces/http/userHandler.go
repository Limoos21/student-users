package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"stud-trener/internal/application"
	"stud-trener/internal/interfaces/dto"
)

type UserHandler struct {
	useCase application.UserUseCaseInterface
	logger  *zap.Logger
}

func NewUserHandler(r *gin.RouterGroup, useCase application.UserUseCaseInterface, logger *zap.Logger) {
	handler := &UserHandler{useCase, logger}
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.AuthorizeUser)
	r.GET("/", handler.GetUserByUsername)
}

func (handler *UserHandler) RegisterUser(c *gin.Context) {
	var userDTO dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		handler.logger.Error("Invalid registration data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err := handler.useCase.RegisterUser(userDTO)
	if err != nil {
		handler.logger.Error("Error registering user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем куки с информацией о пользователе (без пароля)
	c.SetCookie("user_id", strconv.Itoa(int(*userDTO.StudentID)), 3600, "/", c.Request.URL.Hostname(), false, false)
	c.SetCookie("username", userDTO.Username, 3600, "/", c.Request.URL.Hostname(), false, false)
	c.SetCookie("role", userDTO.Role, 3600, "/", c.Request.URL.Hostname(), false, false)

	handler.logger.Info("User registered successfully", zap.Int64("user_id", 32))
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Авторизация пользователя
func (handler *UserHandler) AuthorizeUser(c *gin.Context) {
	var authDTO dto.AuthUserDTO
	if err := c.ShouldBindJSON(&authDTO); err != nil {
		handler.logger.Error("Invalid authorization data", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	userDTO, err := handler.useCase.AuthorizeUser(authDTO)
	if err != nil {
		handler.logger.Error("Authorization failed", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authorization successful",
		"user":    userDTO,
	})
}

// Получение пользователя по имени
func (handler *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username parameter is required"})
		return
	}

	userDTO, err := handler.useCase.GetUserByUsername(username)
	if err != nil {
		handler.logger.Error("Error fetching user by username", zap.String("username", username), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userDTO)
}
