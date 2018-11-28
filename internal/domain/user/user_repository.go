package user

import (
	"context"
	"time"

	"github.com/Jimeux/fked/internal/infra/rdbms"
)

type (
	Repository interface {
		FindOrCreate(ctx context.Context, id ID) (*User, error)
		// UpdateFked(ctx context.Context, userID int64, level int) error
	}

	repository struct {
		rdbms.XormRepository
	}
)

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindOrCreate(ctx context.Context, id ID) (*User, error) {
	u := new(User)
	has, err := r.Session(ctx).ID(id).Get(u)
	if err != nil {
		return nil, err
	}
	if has {
		return u, err
	}

	now := time.Now()
	u.ID = id
	u.Created = now
	u.Updated = now

	if _, err := r.Session(ctx).InsertOne(u); err != nil {
		return nil, err
	}
	return u, nil
}
