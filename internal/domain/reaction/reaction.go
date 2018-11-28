package reaction

import (
	"time"
)

type (
	Code  string
	Level int
)

const (
	_ Level = iota
	FkedUp
	SuperBad
	ReallyBad
	Bad
	NotBad
	Okay
	PrettyGood
	Good
	Great
	Fantastisch
)

type Reaction struct {
	Code      Code      `xorm:"'code' pk(s)"`
	FkedLevel Level     `xorm:"'fked_level'"`
	Created   time.Time `xorm:"'created'"`
}
