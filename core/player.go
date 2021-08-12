package core

import (
	"bitbucket.org/advbet/uid"
)

// Player is a single instance of a connected
// user with abstracted communication channels.
type Player struct {
	Name     string
	GameName string
	Resp     chan Response
}

// NewPlayer returns a new player instance.
func NewPlayer(name string, sign uid.UUID, resp chan Response) *Player {
	return &Player{
		Name: name,
		Resp: resp,
	}
}
