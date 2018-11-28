package fked

import (
	"time"

	"github.com/Jimeux/fked/internal/domain/reaction"
	"github.com/Jimeux/fked/internal/domain/user"
)

type Fked struct {
	UserId       user.ID       `xorm:"'user_id' pk(s)"`
	ReactionCode reaction.Code `xorm:"'reaction_code' pk(s)"`
	Created      time.Time     `xorm:"'created'"`
}
