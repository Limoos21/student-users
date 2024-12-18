package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

type TeamRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTeamRepository(db *gorm.DB, logger *zap.Logger) *TeamRepository {
	return &TeamRepository{db: db, logger: logger}
}

type TeamRepositoryInterface interface {
	GetAll() ([]domain.Team, error)
	GetByID(id int64) (*domain.Team, error)
	DeleteByID(id int64) error
	UpdateByID(id int64, team *domain.Team) error
	Create(team *domain.Team) error
}

// GetAll retrieves all teams from the database.
func (r *TeamRepository) GetAll() ([]domain.Team, error) {
	var teams []domain.Team
	r.logger.Info("Fetching all teams")
	result := r.db.Raw("SELECT * FROM work.team").Scan(&teams)
	if result.Error != nil {
		r.logger.Error("Error fetching teams", zap.Error(result.Error))
		return nil, result.Error
	}
	return teams, nil
}

// GetByID retrieves a single team by ID.
func (r *TeamRepository) GetByID(id int64) (*domain.Team, error) {
	var team domain.Team
	r.logger.Info("Fetching team by ID", zap.Int64("id", id))
	result := r.db.Raw("SELECT * FROM work.team WHERE id = ?", id).Scan(&team)
	if result.Error != nil {
		r.logger.Error("Error fetching team by ID", zap.Int64("id", id), zap.Error(result.Error))
		return nil, result.Error
	}
	return &team, nil
}

// DeleteByID deletes a team by ID.
func (r *TeamRepository) DeleteByID(id int64) error {
	r.logger.Info("Deleting team by ID", zap.Int64("id", id))
	result := r.db.Exec("DELETE FROM work.team WHERE id = ?", id)
	if result.Error != nil {
		r.logger.Error("Error deleting team by ID", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// UpdateByID updates a team's data by ID.
func (r *TeamRepository) UpdateByID(id int64, team *domain.Team) error {
	r.logger.Info("Updating team", zap.Int64("id", id))
	result := r.db.Exec(
		"UPDATE work.team SET name = ?, league = ? WHERE id = ?",
		team.Name, team.League, id,
	)
	if result.Error != nil {
		r.logger.Error("Error updating team", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// Create adds a new team to the database.
func (r *TeamRepository) Create(team *domain.Team) error {
	r.logger.Info("Creating new team")
	result := r.db.Exec(
		"INSERT INTO work.team (name, league) VALUES (?, ?)",
		team.Name, team.League,
	)
	if result.Error != nil {
		r.logger.Error("Error creating team", zap.Error(result.Error))
	}
	return result.Error
}
