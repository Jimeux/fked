package fked

import (
	"context"

	"github.com/Jimeux/fked/internal/domain/reaction"
	"github.com/Jimeux/fked/internal/domain/user"
	"github.com/Jimeux/fked/internal/infra/rdbms"
)

type (
	Service interface {
		UpdateFked(user user.ID, reaction reaction.Code) error
	}

	service struct {
		tx       rdbms.Tx
		fkedRepo Repository
		userRepo user.Repository
	}
)

func NewService(tx rdbms.Tx, repo Repository, userRepo user.Repository) Service {
	return &service{tx, repo, userRepo}
}

func (s *service) UpdateFked(userID user.ID, code reaction.Code) error {
	return s.tx.Exec(func(ctx context.Context) error {
		_, err := s.userRepo.FindOrCreate(ctx, userID)
		if err != nil {
			return err
		}

		f := &Fked{UserId: userID, ReactionCode: code}
		return s.fkedRepo.Create(ctx, f)
	})
}
