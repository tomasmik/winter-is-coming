package core

import (
	"math/rand"
)

type Zombie struct {
	Name string
	Hits int
	x    int
	y    int
}

var names = []string{"night-king", "snow-prince", "ice-face", "coldy-mcold"}

// NewZombie returns a new zombie object
func NewZombie() *Zombie {
	return &Zombie{
		Name: names[rand.Intn(len(names))],
	}
}
