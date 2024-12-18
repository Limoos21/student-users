package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

type TeamTrainerRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTeamTrainerRepository(db *gorm.DB, logger *zap.Logger) *TeamTrainerRepository {
	return &TeamTrainerRepository{db: db, logger: logger}
}

type TeamTrainerRepositoryInterface interface {
	GetAll() ([]domain.TeamTrainer, error)
	GetByID(id int64) (*domain.TeamTrainer, error)
	DeleteByID(id int64) error
	UpdateByID(id int64, teamTrainer *domain.TeamTrainer) error
	Create(teamTrainer *domain.TeamTrainer) error
}

// GetAll retrieves all team-trainer relationships.
func (r *TeamTrainerRepository) GetAll() ([]domain.TeamTrainer, error) {
	var teamTrainers []domain.TeamTrainer
	r.logger.Info("Fetching all team-trainer relationships")
	result := r.db.Raw("SELECT * FROM work.team_trainer").Scan(&teamTrainers)
	if result.Error != nil {
		r.logger.Error("Error fetching team-trainer relationships", zap.Error(result.Error))
		return nil, result.Error
	}
	return teamTrainers, nil
}

// GetByID retrieves a single team-trainer relationship by ID.
func (r *TeamTrainerRepository) GetByID(id int64) (*domain.TeamTrainer, error) {
	var teamTrainer domain.TeamTrainer
	r.logger.Info("Fetching team-trainer relationship by ID", zap.Int64("id", id))
	result := r.db.Raw("SELECT * FROM work.team_trainer WHERE id = ?", id).Scan(&teamTrainer)
	if result.Error != nil {
		r.logger.Error("Error fetching team-trainer relationship by ID", zap.Int64("id", id), zap.Error(result.Error))
		return nil, result.Error
	}
	return &teamTrainer, nil
}

// DeleteByID deletes a team-trainer relationship by ID.
func (r *TeamTrainerRepository) DeleteByID(id int64) error {
	r.logger.Info("Deleting team-trainer relationship by ID", zap.Int64("id", id))
	result := r.db.Exec("DELETE FROM work.team_trainer WHERE id = ?", id)
	if result.Error != nil {
		r.logger.Error("Error deleting team-trainer relationship by ID", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// UpdateByID updates a team-trainer relationship by ID.
func (r *TeamTrainerRepository) UpdateByID(id int64, teamTrainer *domain.TeamTrainer) error {
	r.logger.Info("Updating team-trainer relationship", zap.Int64("id", id))
	result := r.db.Exec(
		"UPDATE work.team_trainer SET id_team = ?, id_trainer = ? WHERE id = ?",
		teamTrainer.TeamID, teamTrainer.TrainerID, id,
	)
	if result.Error != nil {
		r.logger.Error("Error updating team-trainer relationship", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// Create adds a new team-trainer relationship.
func (r *TeamTrainerRepository) Create(teamTrainer *domain.TeamTrainer) error {
	r.logger.Info("Creating new team-trainer relationship")
	result := r.db.Exec(
		"INSERT INTO work.team_trainer (id_team, id_trainer) VALUES (?, ?)",
		teamTrainer.TeamID, teamTrainer.TrainerID,
	)
	if result.Error != nil {
		r.logger.Error("Error creating team-trainer relationship", zap.Error(result.Error))
	}
	return result.Error
}
