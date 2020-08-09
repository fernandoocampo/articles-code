package repository

import (
	"context"

	"github.com/fernandoocampo/articles-code/usefullog/pkg/domain"
	"github.com/fernandoocampo/articles-code/usefullog/pkg/logging"
)

// dblibrary any database library you are using.
type dblibrary struct {
}

func (d *dblibrary) persist() error {
	return nil
}

// AnyStorage is any database logic implementation.
type AnyStorage struct {
	client dblibrary
	logger logging.Logger
}

// NewAnyStorage creates a new storage
func NewAnyStorage(logger logging.Logger) *AnyStorage {
	logger.Debug("creating any storage repository", logging.Fields{"method": "NewAnyStorage"})
	return &AnyStorage{
		client: dblibrary{},
		logger: logger,
	}
}

// Save creates a new employee and returns error if
// something goes wrong.
func (a *AnyStorage) Save(ctx context.Context, employee domain.Employee) error {
	a.logger.Info("storing employee", logging.Fields{"method": "Save", "new employee": employee})
	err := a.client.persist()
	if err != nil {
		a.logger.Error(
			"somethig goes wrong trying to persist a new employee",
			logging.Fields{
				"error":    err,
				"method":   "Save",
				"employee": employee,
			},
		)
		return err
	}
	a.logger.Info("employee was stored succesfully", logging.Fields{"method": "Save", "employee": employee})
	return nil
}
