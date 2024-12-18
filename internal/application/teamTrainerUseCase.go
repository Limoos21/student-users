package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type TeamTrainerUseCase struct {
	Repository repository.TeamTrainerRepositoryInterface
	logger     *zap.Logger
}

type TeamTrainerUseCaseInterface interface {
	CreateTeamTrainer(dto dto.TeamTrainerDTO) (dto.TeamTrainerDTO, error)
	GetAllTeamTrainers() ([]dto.TeamTrainerDTO, error)
	GetTeamTrainerByID(id int64) (dto.TeamTrainerDTO, error)
	UpdateTeamTrainer(id int64, dto dto.TeamTrainerDTO) (dto.TeamTrainerDTO, error)
	DeleteTeamTrainer(id int64) error
}

func NewTeamTrainerUseCase(repository repository.TeamTrainerRepositoryInterface, logger *zap.Logger) *TeamTrainerUseCase {
	return &TeamTrainerUseCase{Repository: repository, logger: logger}
}

// CreateTeamTrainer creates a new team-trainer relationship.
func (t TeamTrainerUseCase) CreateTeamTrainer(dto dto.TeamTrainerDTO) (dto.TeamTrainerDTO, error) {
	newTeamTrainerDomain := domain.TeamTrainer{
		TeamID:    dto.TeamID,
		TrainerID: dto.TrainerID,
	}
	err := t.Repository.Create(&newTeamTrainerDomain)
	if err != nil {
		t.logger.Error("Error creating team-trainer relationship", zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// GetAllTeamTrainers retrieves all team-trainer relationships.
func (t TeamTrainerUseCase) GetAllTeamTrainers() ([]dto.TeamTrainerDTO, error) {
	teamTrainers, err := t.Repository.GetAll()
	if err != nil {
		t.logger.Error("Error fetching all team-trainer relationships", zap.Error(err))
		return nil, err
	}

	newTeamTrainers := make([]dto.TeamTrainerDTO, len(teamTrainers))
	for i, teamTrainer := range teamTrainers {
		newTeamTrainers[i] = dto.TeamTrainerDTO{
			ID:        &teamTrainer.ID,
			TeamID:    teamTrainer.TeamID,
			TrainerID: teamTrainer.TrainerID,
		}
	}
	return newTeamTrainers, nil
}

// GetTeamTrainerByID retrieves a team-trainer relationship by ID.
func (t TeamTrainerUseCase) GetTeamTrainerByID(id int64) (dto.TeamTrainerDTO, error) {
	teamTrainer, err := t.Repository.GetByID(id)
	if err != nil {
		t.logger.Error("Error fetching team-trainer relationship by ID", zap.Int64("id", id), zap.Error(err))
		return dto.TeamTrainerDTO{}, err
	}

	return dto.TeamTrainerDTO{
		ID:        &teamTrainer.ID,
		TeamID:    teamTrainer.TeamID,
		TrainerID: teamTrainer.TrainerID,
	}, nil
}

// UpdateTeamTrainer updates a team-trainer relationship by ID.
func (t TeamTrainerUseCase) UpdateTeamTrainer(id int64, dto dto.TeamTrainerDTO) (dto.TeamTrainerDTO, error) {
	teamTrainerToUpdate := domain.TeamTrainer{
		TeamID:    dto.TeamID,
		TrainerID: dto.TrainerID,
	}
	err := t.Repository.UpdateByID(id, &teamTrainerToUpdate)
	if err != nil {
		t.logger.Error("Error updating team-trainer relationship", zap.Int64("id", id), zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// DeleteTeamTrainer deletes a team-trainer relationship by ID.
func (t TeamTrainerUseCase) DeleteTeamTrainer(id int64) error {
	err := t.Repository.DeleteByID(id)
	if err != nil {
		t.logger.Error("Error deleting team-trainer relationship", zap.Int64("id", id), zap.Error(err))
	}
	return err
}
