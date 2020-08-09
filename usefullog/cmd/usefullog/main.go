package main

import (
	"context"
	"os"

	"github.com/fernandoocampo/articles-code/usefullog/pkg/application"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/domain"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/logging"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/repository"
)

func main() {
	logger := logging.New(logging.Logrus, "usefullog.main", logging.Debug)
	logger.Info("starting usefullog application", nil)

	repositoryLogger := logging.New(logging.Logrus, "repository.AnyStorage", logging.Debug)
	repository := repository.NewAnyStorage(repositoryLogger)
	employeeLogger := logging.New(logging.Logrus, "application.EmployeeService", logging.Debug)
	service := application.NewEmployeeService(repository, employeeLogger)

	newEmployee := domain.NewEmployee{
		FirstName: "Fernando",
		LastName:  "Ocampo",
		Email:     "fernando.ocampo@dev.to",
	}

	ctx := context.Background()
	employeeID, err := service.Create(ctx, newEmployee)
	if err != nil {
		logger.Error("new employee cannot be created", logging.Fields{"error": err})
		os.Exit(1)
	}
	logger.Info("employee was created successfully", logging.Fields{"id": employeeID})

}
