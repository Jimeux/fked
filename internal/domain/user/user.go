package user

import (
	"errors"
	"time"
)

type ID string

type User struct {
	ID      ID        `xorm:"'id' pk"`
	Created time.Time `xorm:"'created'"`
	Updated time.Time `xorm:"'updated'"`
}

func Conv(res interface{}) (*User, error) {
	u, ok := res.(*User)
	if !ok {
		return nil, errors.New("type other than *User passed to user.Conv")
	}
	return u, nil
}
