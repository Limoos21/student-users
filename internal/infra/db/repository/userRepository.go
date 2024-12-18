package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

// Репозиторий для работы с пользователями
type UserRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// Создание нового репозитория
func NewUserRepository(db *gorm.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

// Интерфейс репозитория пользователей
type UserRepositoryInterface interface {
	RegisterUser(user *domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
	AuthorizeUser(username, passwordHash string) (*domain.User, error)
}

// Регистрация пользователя
func (r *UserRepository) RegisterUser(user *domain.User) error {
	r.logger.Info("Registering new user", zap.String("username", user.Username))
	query := `
		INSERT INTO work.users (username, password_hash, role, student_id, trainer_id)
		VALUES (?, ?, ?, ?, ?)`
	result := r.db.Exec(query, user.Username, user.PasswordHash, user.Role, user.StudentID, user.TrainerID)
	if result.Error != nil {
		r.logger.Error("Error registering user", zap.String("username", user.Username), zap.Error(result.Error))
		return result.Error
	}
	return nil
}

// Получение пользователя по имени
func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	r.logger.Info("Fetching user by username", zap.String("username", username))
	query := `
		SELECT id, username, password_hash, role, student_id, trainer_id
		FROM work.users
		WHERE username = ?`
	result := r.db.Raw(query, username).Scan(&user)
	if result.Error != nil {
		r.logger.Error("Error fetching user by username", zap.String("username", username), zap.Error(result.Error))
		return nil, result.Error
	}
	return &user, nil
}

// Авторизация пользователя
func (r *UserRepository) AuthorizeUser(username, passwordHash string) (*domain.User, error) {
	var user domain.User
	r.logger.Info("Authorizing user", zap.String("username", username))
	query := `
		SELECT id, username, password_hash, role, student_id, trainer_id
		FROM work.users
		WHERE username = ? AND password_hash = ?`
	result := r.db.Raw(query, username, passwordHash).Scan(&user)
	if result.Error != nil {
		r.logger.Error("Authorization failed", zap.String("username", username), zap.Error(result.Error))
		return nil, result.Error
	}
	return &user, nil
}
