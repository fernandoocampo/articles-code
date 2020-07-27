package application

import (
	"context"

	"github.com/fernandoocampo/articles-code/usefullog/pkg/domain"
)

// Repository defines behavior
type Repository interface {
	// Save creates a new employee and returns error if
	// something goes wrong.
	Save(ctx context.Context, employee domain.Employee) error
}
