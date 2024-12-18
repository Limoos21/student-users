package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

type TrainerRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTrainerRepository(db *gorm.DB, logger *zap.Logger) *TrainerRepository {
	return &TrainerRepository{db: db, logger: logger}
}

type TrainerRepositoryInterface interface {
	GetAll() ([]domain.Trainer, error)
	GetByID(id int64) (*domain.Trainer, error)
	DeleteByID(id int64) error
	UpdateByID(id int64, trainer *domain.Trainer) error
	Create(trainer *domain.Trainer) error
}

// GetAll retrieves all trainers from the database.
func (r *TrainerRepository) GetAll() ([]domain.Trainer, error) {
	var trainers []domain.Trainer
	r.logger.Info("Fetching all trainers")
	result := r.db.Raw("SELECT * FROM work.trainer").Scan(&trainers)
	if result.Error != nil {
		r.logger.Error("Error fetching trainers", zap.Error(result.Error))
		return nil, result.Error
	}
	return trainers, nil
}

// GetByID retrieves a single trainer by ID.
func (r *TrainerRepository) GetByID(id int64) (*domain.Trainer, error) {
	var trainer domain.Trainer
	r.logger.Info("Fetching trainer by ID", zap.Int64("id", id))
	result := r.db.Raw("SELECT * FROM work.trainer WHERE id = ?", id).Scan(&trainer)
	if result.Error != nil {
		r.logger.Error("Error fetching trainer by ID", zap.Int64("id", id), zap.Error(result.Error))
		return nil, result.Error
	}
	return &trainer, nil
}

// DeleteByID deletes a trainer by ID.
func (r *TrainerRepository) DeleteByID(id int64) error {
	r.logger.Info("Deleting trainer by ID", zap.Int64("id", id))
	result := r.db.Exec("DELETE FROM work.trainer WHERE id = ?", id)
	if result.Error != nil {
		r.logger.Error("Error deleting trainer by ID", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// UpdateByID updates a trainer's data by ID.
func (r *TrainerRepository) UpdateByID(id int64, trainer *domain.Trainer) error {
	r.logger.Info("Updating trainer", zap.Int64("id", id))
	result := r.db.Exec(
		"UPDATE work.trainer SET name = ?, age = ? WHERE id = ?",
		trainer.Name, trainer.Age, id,
	)
	if result.Error != nil {
		r.logger.Error("Error updating trainer", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// Create adds a new trainer to the database.
func (r *TrainerRepository) Create(trainer *domain.Trainer) error {
	r.logger.Info("Creating new trainer")
	result := r.db.Exec(
		"INSERT INTO work.trainer (name, age) VALUES (?, ?)",
		trainer.Name, trainer.Age,
	)
	if result.Error != nil {
		r.logger.Error("Error creating trainer", zap.Error(result.Error))
	}
	return result.Error
}
