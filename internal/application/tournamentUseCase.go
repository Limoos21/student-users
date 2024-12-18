package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type TournamentUseCase struct {
	Repository repository.TournamentRepositoryInterface
	logger     *zap.Logger
}

type TournamentUseCaseInterface interface {
	CreateTournament(dto dto.TournamentDTO) (dto.TournamentDTO, error)
	GetAllTournaments() ([]dto.TournamentDTO, error)
	GetTournamentByID(id int64) (dto.TournamentDTO, error)
	UpdateTournament(id int64, dto dto.TournamentDTO) (dto.TournamentDTO, error)
	DeleteTournament(id int64) error
}

func NewTournamentUseCase(repository repository.TournamentRepositoryInterface, logger *zap.Logger) *TournamentUseCase {
	return &TournamentUseCase{Repository: repository, logger: logger}
}

// CreateTournament creates a new tournament.
func (t TournamentUseCase) CreateTournament(dto dto.TournamentDTO) (dto.TournamentDTO, error) {
	newTournamentDomain := domain.Tournament{
		Name:     dto.Name,
		Room:     dto.Room,
		Datetime: dto.Datetime,
		TeamID:   dto.TeamID,
	}
	err := t.Repository.Create(&newTournamentDomain)
	if err != nil {
		t.logger.Error("Error creating tournament", zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// GetAllTournaments retrieves all tournaments.
func (t TournamentUseCase) GetAllTournaments() ([]dto.TournamentDTO, error) {
	tournaments, err := t.Repository.GetAll()
	if err != nil {
		t.logger.Error("Error fetching all tournaments", zap.Error(err))
		return nil, err
	}

	newTournaments := make([]dto.TournamentDTO, len(tournaments))
	for i, tournament := range tournaments {
		newTournaments[i] = dto.TournamentDTO{
			ID:       &tournament.ID,
			Name:     tournament.Name,
			Room:     tournament.Room,
			Datetime: tournament.Datetime,
			TeamID:   tournament.TeamID,
		}
	}
	return newTournaments, nil
}

// GetTournamentByID retrieves a tournament by ID.
func (t TournamentUseCase) GetTournamentByID(id int64) (dto.TournamentDTO, error) {
	tournament, err := t.Repository.GetByID(id)
	if err != nil {
		t.logger.Error("Error fetching tournament by ID", zap.Int64("id", id), zap.Error(err))
		return dto.TournamentDTO{}, err
	}

	return dto.TournamentDTO{
		ID:       &tournament.ID,
		Name:     tournament.Name,
		Room:     tournament.Room,
		Datetime: tournament.Datetime,
		TeamID:   tournament.TeamID,
	}, nil
}

// UpdateTournament updates a tournament by ID.
func (t TournamentUseCase) UpdateTournament(id int64, dto dto.TournamentDTO) (dto.TournamentDTO, error) {
	tournamentToUpdate := domain.Tournament{
		Name:     dto.Name,
		Room:     dto.Room,
		Datetime: dto.Datetime,
		TeamID:   dto.TeamID,
	}
	err := t.Repository.UpdateByID(id, &tournamentToUpdate)
	if err != nil {
		t.logger.Error("Error updating tournament", zap.Int64("id", id), zap.Error(err))
		return dto, err
	}
	return dto, nil
}

// DeleteTournament deletes a tournament by ID.
func (t TournamentUseCase) DeleteTournament(id int64) error {
	err := t.Repository.DeleteByID(id)
	if err != nil {
		t.logger.Error("Error deleting tournament", zap.Int64("id", id), zap.Error(err))
	}
	return err
}
