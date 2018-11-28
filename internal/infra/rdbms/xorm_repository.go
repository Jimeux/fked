package rdbms

import (
	"context"
	"fmt"

	"github.com/go-xorm/xorm"
)

type XormRepository struct {
}

// Session returns a xorm session (xorm.Interface) instance contained in ctx.
// DANGER! It will panic if ctx contains no session instance!
func (r *XormRepository) Session(ctx context.Context) xorm.Interface {
	session, ok := ctx.Value(SessionKey).(xorm.Interface)
	if !ok {
		panic("database session not passed to repository method")
	}
	return session
}

func (r *XormRepository) VerifyAffected(expected int64, actual int64, err error) error {
	if err != nil {
		return err
	}
	if expected != actual {
		return fmt.Errorf("expected %d rows to be updated, but actual was %d", expected, actual)
	}
	return nil
}
