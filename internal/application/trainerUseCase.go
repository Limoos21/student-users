package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type TrainerUseCase struct {
	Repository repository.TrainerRepositoryInterface
	logger     *zap.Logger
}

type TrainerUseCaseInterface interface {
	CreateTrainer(dto dto.TrainerDTO) (dto.TrainerDTO, error)
	GetAllTrainers() ([]dto.TrainerDTO, error)
	GetTrainerByID(id int64) (dto.TrainerDTO, error)
	UpdateTrainer(id int64, dto dto.TrainerDTO) (dto.TrainerDTO, error)
	DeleteTrainer(id int64) error
}

func NewTrainerUseCase(repository repository.TrainerRepositoryInterface, logger *zap.Logger) *TrainerUseCase {
	return &TrainerUseCase{Repository: repository, logger: logger}
}

// CreateTrainer creates a new trainer.
func (t TrainerUseCase) CreateTrainer(dto dto.TrainerDTO) (dto.TrainerDTO, error) {
	newTrainerDomain := domain.Trainer{
		Name: dto.Name,
		Age:  dto.Age,
	}
	err := t.Repository.Create(&newTrainerDomain)
	if err != nil {
		t.logger.Error("Error creating trainer", zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// GetAllTrainers retrieves all trainers.
func (t TrainerUseCase) GetAllTrainers() ([]dto.TrainerDTO, error) {
	trainers, err := t.Repository.GetAll()
	if err != nil {
		t.logger.Error("Error fetching all trainers", zap.Error(err))
		return nil, err
	}

	newTrainers := make([]dto.TrainerDTO, len(trainers))
	for i, trainer := range trainers {
		newTrainers[i] = dto.TrainerDTO{
			ID:   &trainer.ID,
			Name: trainer.Name,
			Age:  trainer.Age,
		}
	}
	return newTrainers, nil
}

// GetTrainerByID retrieves a trainer by ID.
func (t TrainerUseCase) GetTrainerByID(id int64) (dto.TrainerDTO, error) {
	trainer, err := t.Repository.GetByID(id)
	if err != nil {
		t.logger.Error("Error fetching trainer by ID", zap.Int64("id", id), zap.Error(err))
		return dto.TrainerDTO{}, err
	}

	return dto.TrainerDTO{
		ID:   &trainer.ID,
		Name: trainer.Name,
		Age:  trainer.Age,
	}, nil
}

// UpdateTrainer updates a trainer by ID.
func (t TrainerUseCase) UpdateTrainer(id int64, dto dto.TrainerDTO) (dto.TrainerDTO, error) {
	trainerToUpdate := domain.Trainer{
		Name: dto.Name,
		Age:  dto.Age,
	}
	err := t.Repository.UpdateByID(id, &trainerToUpdate)
	if err != nil {
		t.logger.Error("Error updating trainer", zap.Int64("id", id), zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// DeleteTrainer deletes a trainer by ID.
func (t TrainerUseCase) DeleteTrainer(id int64) error {
	err := t.Repository.DeleteByID(id)
	if err != nil {
		t.logger.Error("Error deleting trainer", zap.Int64("id", id), zap.Error(err))
	}
	return err
}
