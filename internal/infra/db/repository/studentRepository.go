package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"stud-trener/internal/domain"
)

type StudentRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewStudentRepository(db *gorm.DB, logger *zap.Logger) *StudentRepository {
	return &StudentRepository{db: db, logger: logger}
}

type StudentRepositoryInterface interface {
	GetAll() ([]domain.Student, error)
	GetByID(id int64) (*domain.Student, error)
	DeleteByID(id int64) error
	UpdateByID(id int64, student *domain.Student) error
	Create(student *domain.Student) error
}

// GetAll retrieves all students from the database.
func (r *StudentRepository) GetAll() ([]domain.Student, error) {
	var students []domain.Student
	r.logger.Info("Fetching all students")
	result := r.db.Raw("SELECT * FROM work.student").Scan(&students)
	if result.Error != nil {
		r.logger.Error("Error fetching students", zap.Error(result.Error))
		return nil, result.Error
	}
	return students, nil
}

// GetByID retrieves a single student by ID.
func (r *StudentRepository) GetByID(id int64) (*domain.Student, error) {
	var student domain.Student
	r.logger.Info("Fetching student by ID", zap.Int64("id", id))
	result := r.db.Raw("SELECT * FROM work.student WHERE id = ?", id).Scan(&student)
	if result.Error != nil {
		r.logger.Error("Error fetching student by ID", zap.Int64("id", id), zap.Error(result.Error))
		return nil, result.Error
	}
	return &student, nil
}

// DeleteByID deletes a student by ID.
func (r *StudentRepository) DeleteByID(id int64) error {
	r.logger.Info("Deleting student by ID", zap.Int64("id", id))
	result := r.db.Exec("DELETE FROM work.student WHERE id = ?", id)
	if result.Error != nil {
		r.logger.Error("Error deleting student by ID", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// UpdateByID updates a student's data by ID.
func (r *StudentRepository) UpdateByID(id int64, student *domain.Student) error {
	r.logger.Info("Updating student", zap.Int64("id", id))
	result := r.db.Exec(
		"UPDATE work.student SET name = ?, age = ?, height = ?, weight = ?, team_id = ? WHERE id = ?",
		student.Name, student.Age, student.Height, student.Weight, student.TeamID, id,
	)
	if result.Error != nil {
		r.logger.Error("Error updating student", zap.Int64("id", id), zap.Error(result.Error))
	}
	return result.Error
}

// Create adds a new student to the database.
func (r *StudentRepository) Create(student *domain.Student) error {
	r.logger.Info("Creating new student", zap.String("name", student.Name))
	result := r.db.Exec(
		"INSERT INTO work.student (name, age, height, weight, team_id) VALUES (?, ?, ?, ?, ?)",
		student.Name, student.Age, student.Height, student.Weight, student.TeamID,
	)
	if result.Error != nil {
		r.logger.Error("Error creating student", zap.String("name", student.Name), zap.Error(result.Error))
	}
	return result.Error
}
