package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

type TrainRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTrainRepository(db *gorm.DB, logger *zap.Logger) *TrainRepository {
	return &TrainRepository{db: db, logger: logger}
}

type TrainRepositoryInterface interface {
	GetAll() ([]domain.Train, error)
	GetByID(id int64) (*domain.Train, error)
	DeleteByID(id int64) error
	UpdateByID(id int64, train *domain.Train) error
	Create(train *domain.Train) error
}

// GetAll retrieves all training sessions from the database.
func (r *TrainRepository) GetAll() ([]domain.Train, error) {
	var trains []domain.Train
	r.logger.Info("Fetching all training sessions")
	result := r.db.Raw("SELECT * FROM work.train").Scan(&trains)
	if result.Error != nil {
		r.logger.Error("Error fetching training sessions", zap.Error(result.Error))
		return nil, result.Error
	}
	return trains, nil
}

// GetByID retrieves a single training session by ID.
func (r *TrainRepository) GetByID(id int64) (*domain.Train, error) {
	var train domain.Train
	r.logger.Info("Fetching training session by ID", zap.Int64("id", id))
	result := r.db.Raw("SELECT * FROM work.train WHERE id = ?", id).Scan(&train)
	if result.Error != nil {
		r.logger.Error("Error fetching training session by ID", zap.Int64("id", id), zap.Error(result.Error))
		return nil, result.Error
	}
	return &train, nil
}

// DeleteByID deletes a training session by ID.
func (r *TrainRepository) DeleteByID(id int64) error {
	r.logger.Info("Deleting training session by ID", zap.Int64("id", id))
	result := r.db.Exec("DELETE FROM work.train WHERE id = ?", id)
	if result.Error != nil {
		r.logger.Error("Error deleting training session by ID", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// UpdateByID updates a training session by ID.
func (r *TrainRepository) UpdateByID(id int64, train *domain.Train) error {
	r.logger.Info("Updating training session", zap.Int64("id", id))
	result := r.db.Exec(
		"UPDATE work.train SET type = ?, room = ?, datetime = ?, id_trainer = ?, id_team = ? WHERE id = ?",
		train.Type, train.Room, train.Datetime, train.TrainerID, train.TeamID, id,
	)
	if result.Error != nil {
		r.logger.Error("Error updating training session", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// Create adds a new training session to the database.
func (r *TrainRepository) Create(train *domain.Train) error {
	r.logger.Info("Creating new training session")
	result := r.db.Exec(
		"INSERT INTO work.train (type, room, datetime, id_trainer, id_team) VALUES (?, ?, ?, ?, ?)",
		train.Type, train.Room, train.Datetime, train.TrainerID, train.TeamID,
	)
	if result.Error != nil {
		r.logger.Error("Error creating training session", zap.Error(result.Error))
	}
	return result.Error
}
