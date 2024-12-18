package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type TrainUseCase struct {
	Repository repository.TrainRepositoryInterface
	logger     *zap.Logger
}

type TrainUseCaseInterface interface {
	CreateTrain(dto dto.TrainDTO) (dto.TrainDTO, error)
	GetAllTrains() ([]dto.TrainDTO, error)
	GetTrainByID(id int64) (dto.TrainDTO, error)
	UpdateTrain(id int64, dto dto.TrainDTO) (dto.TrainDTO, error)
	DeleteTrain(id int64) error
}

func NewTrainUseCase(repository repository.TrainRepositoryInterface, logger *zap.Logger) *TrainUseCase {
	return &TrainUseCase{Repository: repository, logger: logger}
}

// CreateTrain creates a new training session.
func (t TrainUseCase) CreateTrain(dto dto.TrainDTO) (dto.TrainDTO, error) {
	newTrainDomain := domain.Train{
		Type:      dto.Type,
		Room:      dto.Room,
		Datetime:  dto.Datetime,
		TrainerID: dto.TrainerID,
		TeamID:    dto.TeamID,
	}
	err := t.Repository.Create(&newTrainDomain)
	if err != nil {
		t.logger.Error("Error creating training session", zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// GetAllTrains retrieves all training sessions.
func (t TrainUseCase) GetAllTrains() ([]dto.TrainDTO, error) {
	trains, err := t.Repository.GetAll()
	if err != nil {
		t.logger.Error("Error fetching all training sessions", zap.Error(err))
		return nil, err
	}

	newTrains := make([]dto.TrainDTO, len(trains))
	for i, train := range trains {
		newTrains[i] = dto.TrainDTO{
			ID:        &train.ID,
			Type:      train.Type,
			Room:      train.Room,
			Datetime:  train.Datetime,
			TrainerID: train.TrainerID,
			TeamID:    train.TeamID,
		}
	}
	return newTrains, nil
}

// GetTrainByID retrieves a training session by ID.
func (t TrainUseCase) GetTrainByID(id int64) (dto.TrainDTO, error) {
	train, err := t.Repository.GetByID(id)
	if err != nil {
		t.logger.Error("Error fetching training session by ID", zap.Int64("id", id), zap.Error(err))
		return dto.TrainDTO{}, err
	}

	return dto.TrainDTO{
		ID:        &train.ID,
		Type:      train.Type,
		Room:      train.Room,
		Datetime:  train.Datetime,
		TrainerID: train.TrainerID,
		TeamID:    train.TeamID,
	}, nil
}

// UpdateTrain updates a training session by ID.
func (t TrainUseCase) UpdateTrain(id int64, dto dto.TrainDTO) (dto.TrainDTO, error) {
	trainToUpdate := domain.Train{
		Type:      dto.Type,
		Room:      dto.Room,
		Datetime:  dto.Datetime,
		TrainerID: dto.TrainerID,
		TeamID:    dto.TeamID,
	}
	err := t.Repository.UpdateByID(id, &trainToUpdate)
	if err != nil {
		t.logger.Error("Error updating training session", zap.Int64("id", id), zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// DeleteTrain deletes a training session by ID.
func (t TrainUseCase) DeleteTrain(id int64) error {
	err := t.Repository.DeleteByID(id)
	if err != nil {
		t.logger.Error("Error deleting training session", zap.Int64("id", id), zap.Error(err))
	}
	return err
}
