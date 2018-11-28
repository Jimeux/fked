package fked

import (
	"context"
	"time"

	"github.com/Jimeux/fked/internal/infra/rdbms"
)

type (
	Repository interface {
		Create(ctx context.Context, f *Fked) error
	}

	repository struct {
		rdbms.XormRepository
	}
)

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(ctx context.Context, f *Fked) error {
	f.Created = time.Now()
	rows, err := r.Session(ctx).InsertOne(f)
	return r.VerifyAffected(1, rows, err)
}
