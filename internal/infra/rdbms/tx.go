package rdbms

import (
	"context"
	"fmt"

	"github.com/go-xorm/xorm"
)

const (
	SessionKey = "session_key_akl23j4l_k2j34lkj"
)

type (
	Session  = context.Context
	TxUnit   = func(ctx context.Context) error
	TxResult = func(ctx context.Context) (interface{}, error)

	Tx interface {
		// AutoSession returns an auto-committing session instance.
		AutoSession() Session

		// Exec performs a transaction with no result.
		// It will auto-rollback on error.
		Exec(work TxUnit) error

		// Result performs a transaction with an untyped result.
		// It will auto-rollback on error.
		Result(work TxResult) (interface{}, error)
	}

	tx struct {
		db xorm.EngineInterface
	}
)

func NewTx(db xorm.EngineInterface) Tx {
	return &tx{db}
}

func (t *tx) AutoSession() Session {
	return context.WithValue(context.Background(), SessionKey, t.db)
}

func (t *tx) Exec(work TxUnit) error {
	session := t.db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	if err := work(sessionContext(session)); err != nil {
		if rollErr := session.Rollback(); rollErr != nil {
			return fmt.Errorf("%v: %v", rollErr, err)
		}
		return err
	}
	if err := session.Commit(); err != nil {
		return err
	}
	return nil
}

func (t *tx) Result(work TxResult) (interface{}, error) {
	session := t.db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return nil, err
	}
	ret, err := work(sessionContext(session))
	if err != nil {
		if rollErr := session.Rollback(); rollErr != nil {
			return nil, fmt.Errorf("%v: %v", rollErr, err)
		}
		return nil, err
	}
	if err := session.Commit(); err != nil {
		return nil, err
	}
	return ret, nil
}

func sessionContext(session xorm.Interface) Session {
	return context.WithValue(context.Background(), SessionKey, session)
}
