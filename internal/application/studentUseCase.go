package application

import (
	"go.uber.org/zap"
	"stud-trener/internal/domain"
	"stud-trener/internal/infra/db/repository"
	"stud-trener/internal/interfaces/dto"
)

type StudentUseCase struct {
	Repository repository.StudentRepositoryInterface
	logger     *zap.Logger
}

type StudentUseCaseInterface interface {
	CreateStudent(dto dto.StudentDTO) (dto.StudentDTO, error)
	GetAllStudent() ([]dto.StudentDTO, error)
	GetStudentById(id int) (dto.StudentDTO, error)
	UpdateStudent(id int, dto dto.StudentDTO) (dto.StudentDTO, error)
	DeleteStudent(id int) error
}

func NewStudentUseCase(repository repository.StudentRepositoryInterface, logger *zap.Logger) *StudentUseCase {
	return &StudentUseCase{Repository: repository, logger: logger}
}

func (s *StudentUseCase) CreateStudent(dto dto.StudentDTO) (dto.StudentDTO, error) {
	// Преобразуем DTO в доменный объект
	newStudentDomain := domain.Student{
		Name:   dto.Name,
		Age:    dto.Age,
		Height: dto.Height,
		Weight: dto.Weight,
		TeamID: dto.TeamID,
	}

	// Вызываем метод репозитория для создания студента
	err := s.Repository.Create(&newStudentDomain)
	if err != nil {
		return dto, err
	}

	// Возвращаем DTO после создания
	dto.ID = &newStudentDomain.ID
	return dto, nil
}

func (s *StudentUseCase) GetAllStudent() ([]dto.StudentDTO, error) {
	students, err := s.Repository.GetAll()
	if err != nil {
		return []dto.StudentDTO{}, err
	}

	// Преобразуем список студентов из доменной модели в DTO
	newStudents := make([]dto.StudentDTO, len(students))
	for i, student := range students {
		newStudents[i] = dto.StudentDTO{
			ID:     &student.ID,
			Name:   student.Name,
			Age:    student.Age,
			Height: student.Height,
			Weight: student.Weight,
			TeamID: student.TeamID,
		}
	}
	return newStudents, nil
}

func (s *StudentUseCase) GetStudentById(id int) (dto.StudentDTO, error) {
	student, err := s.Repository.GetByID(int64(id))
	if err != nil {
		return dto.StudentDTO{}, err
	}

	// Преобразуем доменный объект в DTO
	return dto.StudentDTO{
		ID:     &student.ID,
		Name:   student.Name,
		Age:    student.Age,
		Height: student.Height,
		Weight: student.Weight,
		TeamID: student.TeamID,
	}, nil
}

func (s *StudentUseCase) UpdateStudent(id int, dto dto.StudentDTO) (dto.StudentDTO, error) {
	// Преобразуем DTO в доменный объект
	studentDomain := domain.Student{
		Name:   dto.Name,
		Age:    dto.Age,
		Height: dto.Height,
		Weight: dto.Weight,
		TeamID: dto.TeamID,
	}

	// Вызываем репозиторий для обновления студента
	err := s.Repository.UpdateByID(int64(id), &studentDomain)
	if err != nil {
		return dto, err
	}

	// Возвращаем обновленный DTO
	dto.ID = &id
	return dto, nil
}

func (s *StudentUseCase) DeleteStudent(id int) error {
	// Вызываем репозиторий для удаления студента
	err := s.Repository.DeleteByID(int64(id))
	if err != nil {
		return err
	}
	return nil
}
