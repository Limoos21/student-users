package application

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type UserUseCase struct {
	Repository repository.UserRepositoryInterface
	logger     *zap.Logger
}

type UserUseCaseInterface interface {
	RegisterUser(userDTO dto.RegisterUserDTO) error
	AuthorizeUser(authDTO dto.AuthUserDTO) (*dto.UserDTO, error)
	GetUserByUsername(username string) (*dto.UserDTO, error)
}

func NewUserUseCase(repository repository.UserRepositoryInterface, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{Repository: repository, logger: logger}
}

// Регистрация пользователя
func (u *UserUseCase) RegisterUser(userDTO dto.RegisterUserDTO) error {
	u.logger.Info("Registering user", zap.String("username", userDTO.Username))

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("Error hashing password", zap.Error(err))
		return err
	}

	// Создание доменной модели пользователя
	user := &domain.User{
		Username:     userDTO.Username,
		PasswordHash: string(hashedPassword),
		Role:         userDTO.Role,
		StudentID:    userDTO.StudentID,
		TrainerID:    userDTO.TrainerID,
	}

	// Регистрация в репозитории
	err = u.Repository.RegisterUser(user)
	if err != nil {
		u.logger.Error("Error registering user in repository", zap.String("username", userDTO.Username), zap.Error(err))
		return err
	}
	return nil
}

// Авторизация пользователя
func (u *UserUseCase) AuthorizeUser(authDTO dto.AuthUserDTO) (*dto.UserDTO, error) {
	u.logger.Info("Authorizing user", zap.String("username", authDTO.Username))

	// Получение пользователя из репозитория
	user, err := u.Repository.GetUserByUsername(authDTO.Username)
	if err != nil {
		u.logger.Error("Error fetching user for authorization", zap.String("username", authDTO.Username), zap.Error(err))
		return nil, err
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(authDTO.Password))
	if err != nil {
		u.logger.Error("Invalid password for user", zap.String("username", authDTO.Username))
		return nil, errors.New("invalid username or password")
	}

	// Преобразуем доменную модель в DTO
	return &dto.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		StudentID: user.StudentID,
		TrainerID: user.TrainerID,
	}, nil
}

// Получение пользователя по имени
func (u *UserUseCase) GetUserByUsername(username string) (*dto.UserDTO, error) {
	u.logger.Info("Fetching user by username", zap.String("username", username))

	// Получение пользователя из репозитория
	user, err := u.Repository.GetUserByUsername(username)
	if err != nil {
		u.logger.Error("Error fetching user", zap.String("username", username), zap.Error(err))
		return nil, err
	}

	// Преобразуем доменную модель в DTO
	return &dto.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		StudentID: user.StudentID,
		TrainerID: user.TrainerID,
	}, nil
}
