package application

import (
	"context"

	"github.com/fernandoocampo/articles-code/usefullog/pkg/domain"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/logging"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// EmployeeService defines basic logic to store employee data.
type EmployeeService struct {
	logger     logging.Logger
	repository Repository
}

// NewEmployeeService creates a new EmployeeService instance.
func NewEmployeeService(repository Repository, logger logging.Logger) *EmployeeService {
	logger.Debug("creating employee service", logging.Fields{"method": "NewEmployeeService"})
	return &EmployeeService{
		logger:     logger,
		repository: repository,
	}
}

// Create creates a new employee with the given data.
// Returns the ID generated or an error if something goes wrong.
func (e *EmployeeService) Create(ctx context.Context, newEmployee domain.NewEmployee) (string, error) {
	e.logger.Debug("creating employee", logging.Fields{"method": "Create", "new employee": newEmployee})
	validationError := newEmployee.Validate()
	if validationError != nil {
		return "", errors.Wrap(validationError, "the data of the new employee is not valid")
	}
	newEmployeeID := uuid.New().String()
	e.logger.Debug("new employee id is generated", logging.Fields{"method": "Create", "id": newEmployeeID})
	employee := newEmployee.ToEmployee(newEmployeeID)
	e.logger.Info("storing employee", logging.Fields{"method": "Create", "employee": employee})
	err := e.repository.Save(ctx, employee)
	if err != nil {
		e.logger.Error(
			"employee cannot be stored",
			logging.Fields{
				"method":   "Create",
				"error":    err,
				"employee": employee,
			},
		)
		return "", errors.New("employee cannot be stored")
	}
	return newEmployeeID, nil
}
