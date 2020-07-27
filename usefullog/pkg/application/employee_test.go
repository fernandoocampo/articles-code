package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fernandoocampo/articles-code/usefullog/pkg/application"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/domain"
)

func TestNewEmployee(t *testing.T) {
	// GIVEN
	newEmployee := domain.NewEmployee{
		FirstName: "William",
		LastName:  "Wallace",
		Email:     "william.wallace@scotland.com",
	}
	repository := newRepositoryMock(nil)
	service := application.NewEmployeeService(repository)
	// WHEN
	newID, err := service.Create(context.TODO(), newEmployee)
	// THEN
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(newID) < 1 {
		t.Errorf("empty ID is not expected")
	}
}

func TestNewEmployeeUnexpectedError(t *testing.T) {
	// GIVEN
	newEmployee := domain.NewEmployee{
		FirstName: "William",
		LastName:  "Wallace",
		Email:     "william.wallace@scotland.com",
	}
	repository := newRepositoryMock(errors.New("error"))
	service := application.NewEmployeeService(repository)
	// WHEN
	newID, err := service.Create(context.TODO(), newEmployee)
	// THEN
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(newID) < 1 {
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
