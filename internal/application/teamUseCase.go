package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type TeamUseCase struct {
	Repository repository.TeamRepositoryInterface
	logger     *zap.Logger
}

type TeamUseCaseInterface interface {
	CreateTeam(dto dto.TeamDTO) (dto.TeamDTO, error)
	GetAllTeams() ([]dto.TeamDTO, error)
	GetTeamByID(id int) (dto.TeamDTO, error)
	UpdateTeam(id int, dto dto.TeamDTO) (dto.TeamDTO, error)
	DeleteTeam(id int) error
}

func NewTeamUseCase(repository repository.TeamRepositoryInterface, logger *zap.Logger) *TeamUseCase {
	return &TeamUseCase{Repository: repository, logger: logger}
}

// CreateTeam creates a new team.
func (t TeamUseCase) CreateTeam(dto dto.TeamDTO) (dto.TeamDTO, error) {
	newTeamDomain := domain.Team{
		Name:   dto.Name,
		League: dto.League,
	}
	err := t.Repository.Create(&newTeamDomain)
	if err != nil {
		t.logger.Error("Error creating team", zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// GetAllTeams retrieves all teams.
func (t TeamUseCase) GetAllTeams() ([]dto.TeamDTO, error) {
	teams, err := t.Repository.GetAll()
	if err != nil {
		t.logger.Error("Error fetching all teams", zap.Error(err))
		return nil, err
	}

	newTeams := make([]dto.TeamDTO, len(teams))
	for i, team := range teams {
		newTeams[i] = dto.TeamDTO{
			ID:     &team.ID,
			Name:   team.Name,
			League: team.League,
		}
	}
	return newTeams, nil
}

// GetTeamByID retrieves a team by ID.
func (t TeamUseCase) GetTeamByID(id int) (dto.TeamDTO, error) {
	team, err := t.Repository.GetByID(int64(id))
	if err != nil {
		t.logger.Error("Error fetching team by ID", zap.Int("id", id), zap.Error(err))
		return dto.TeamDTO{}, err
	}

	return dto.TeamDTO{
		ID:     &team.ID,
		Name:   team.Name,
		League: team.League,
	}, nil
}

// UpdateTeam updates a team's details.
func (t TeamUseCase) UpdateTeam(id int, dto dto.TeamDTO) (dto.TeamDTO, error) {
	teamToUpdate := domain.Team{
		Name:   dto.Name,
		League: dto.League,
	}
	err := t.Repository.UpdateByID(int64(id), &teamToUpdate)
	if err != nil {
		t.logger.Error("Error updating team", zap.Int("id", id), zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// DeleteTeam deletes a team by ID.
func (t TeamUseCase) DeleteTeam(id int) error {
	err := t.Repository.DeleteByID(int64(id))
	if err != nil {
		t.logger.Error("Error deleting team", zap.Int("id", id), zap.Error(err))
	}
	return err
}
