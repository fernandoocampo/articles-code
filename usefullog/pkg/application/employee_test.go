package application_test

import (
	"bytes"
	"context"
	"errors"
	"log"
	"testing"

	"github.com/fernandoocampo/articles-code/usefullog/pkg/application"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/domain"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/logging"
)

func TestNewEmployee(t *testing.T) {
	// GIVEN
	newEmployee := domain.NewEmployee{
		FirstName: "William",
		LastName:  "Wallace",
		Email:     "william.wallace@scotland.com",
	}
	loggerBuffer := new(bytes.Buffer)
	repository := newRepositoryMock(nil)
	employeeLogger := logging.NewWithWriter(logging.Logrus, "application.EmployeeService", logging.Debug, loggerBuffer)
	service := application.NewEmployeeService(repository, employeeLogger)
	// WHEN
	newID, err := service.Create(context.TODO(), newEmployee)
	// THEN
	if err != nil {
		log.Println(loggerBuffer.String())
		t.Errorf("unexpected error: %s", err)
	}
	if len(newID) < 1 {
		log.Println(loggerBuffer.String())
		t.Errorf("empty ID is not expected")
	}
}

func TestNewEmployeeValidationdError(t *testing.T) {
	// GIVEN
	newEmployee := domain.NewEmployee{
		FirstName: "William",
		LastName:  "Wallace",
	}
	loggerBuffer := new(bytes.Buffer)
	repository := newRepositoryMock(nil)
	employeeLogger := logging.NewWithWriter(logging.Logrus, "application.EmployeeService", logging.Debug, loggerBuffer)
	service := application.NewEmployeeService(repository, employeeLogger)
	// WHEN
	newID, err := service.Create(context.TODO(), newEmployee)
	// THEN
	if err != nil {
		log.Println(loggerBuffer.String())
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}
	if len(newID) < 1 {
		log.Println(loggerBuffer.String())
		t.Error("empty ID is not expected")
	}
}

func TestNewEmployeeUnexpectedError(t *testing.T) {
	// GIVEN
	newEmployee := domain.NewEmployee{
		FirstName: "William",
		LastName:  "Wallace",
		Email:     "william.wallace@scotland.com",
	}
	loggerBuffer := new(bytes.Buffer)
	repository := newRepositoryMock(errors.New("any error"))
	employeeLogger := logging.NewWithWriter(logging.Logrus, "application.EmployeeService", logging.Debug, loggerBuffer)
	service := application.NewEmployeeService(repository, employeeLogger)
	// WHEN
	newID, err := service.Create(context.TODO(), newEmployee)
	// THEN
	if err != nil {
		log.Println(loggerBuffer.String())
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}
	if len(newID) < 1 {
		log.Println(loggerBuffer.String())
		t.Errorf("empty ID is not expected")
	}
}

// repositoryMock simulates a repository for employee service.
type repositoryMock struct {
	err error
}

// newRepositoryMock creates a new repository.
func newRepositoryMock(err error) *repositoryMock {
	return &repositoryMock{
		err: err,
	}
}

// Save mocks save behavior
func (r *repositoryMock) Save(ctx context.Context, employee domain.Employee) error {
	if r.err != nil {
		return r.err
	}
	return nil
}
