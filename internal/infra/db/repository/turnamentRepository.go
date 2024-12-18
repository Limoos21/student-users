package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

type TournamentRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTournamentRepository(db *gorm.DB, logger *zap.Logger) *TournamentRepository {
	return &TournamentRepository{db: db, logger: logger}
}

type TournamentRepositoryInterface interface {
	GetAll() ([]domain.Tournament, error)
	GetByID(id int64) (*domain.Tournament, error)
	DeleteByID(id int64) error
	UpdateByID(id int64, tournament *domain.Tournament) error
	Create(tournament *domain.Tournament) error
}

// GetAll retrieves all tournaments from the database.
func (r *TournamentRepository) GetAll() ([]domain.Tournament, error) {
	var tournaments []domain.Tournament
	r.logger.Info("Fetching all tournaments")
	result := r.db.Raw("SELECT * FROM work.tournament").Scan(&tournaments)
	if result.Error != nil {
		r.logger.Error("Error fetching tournaments", zap.Error(result.Error))
		return nil, result.Error
	}
	return tournaments, nil
}

// GetByID retrieves a single tournament by ID.
func (r *TournamentRepository) GetByID(id int64) (*domain.Tournament, error) {
	var tournament domain.Tournament
	r.logger.Info("Fetching tournament by ID", zap.Int64("id", id))
	result := r.db.Raw("SELECT * FROM work.tournament WHERE id = ?", id).Scan(&tournament)
	if result.Error != nil {
		r.logger.Error("Error fetching tournament by ID", zap.Int64("id", id), zap.Error(result.Error))
		return nil, result.Error
	}
	return &tournament, nil
}

// DeleteByID deletes a tournament by ID.
func (r *TournamentRepository) DeleteByID(id int64) error {
	r.logger.Info("Deleting tournament by ID", zap.Int64("id", id))
	result := r.db.Exec("DELETE FROM work.tournament WHERE id = ?", id)
	if result.Error != nil {
		r.logger.Error("Error deleting tournament by ID", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// UpdateByID updates a tournament's data by ID.
func (r *TournamentRepository) UpdateByID(id int64, tournament *domain.Tournament) error {
	r.logger.Info("Updating tournament", zap.Int64("id", id))
	result := r.db.Exec(
		"UPDATE work.tournament SET name = ?, room = ?, datetime = ?, id_team = ? WHERE id = ?",
		tournament.Name, tournament.Room, tournament.Datetime, tournament.TeamID, id,
	)
	if result.Error != nil {
		r.logger.Error("Error updating tournament", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// Create adds a new tournament to the database.
func (r *TournamentRepository) Create(tournament *domain.Tournament) error {
	r.logger.Info("Creating new tournament")
	result := r.db.Exec(
		"INSERT INTO work.tournament (name, room, datetime, id_team) VALUES (?, ?, ?, ?)",
		tournament.Name, tournament.Room, tournament.Datetime, tournament.TeamID,
	)
	if result.Error != nil {
		r.logger.Error("Error creating tournament", zap.Error(result.Error))
	}
	return result.Error
}
