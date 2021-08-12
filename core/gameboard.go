package core

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Gameboard struct {
	Zombie *Zombie
}

var (
	maxX = 10
	maxY = 30

	axiX = "x"
	axiY = "y"
)

var axies = []string{axiX, axiY}

func NewGameBoard() *Gameboard {
	return &Gameboard{
		Zombie: NewZombie(),
	}
}

// ZombieWalk makes the Zombie walk in random direction
// returning current x and y coordinates.
func (g *Gameboard) ZombieWalk() (int, int) {
	axi := rand.Intn(len(axies))
	if axies[axi] == axiX && g.Zombie.x < maxX {
		g.Zombie.x++
	}

	if axies[axi] == axiY && g.Zombie.y < maxY {
		g.Zombie.y++
	}
	return g.Zombie.x, g.Zombie.y
}

// ZombieReachedWall returns true if a Zombie has reached the wall
func (g *Gameboard) ZombieReachedWall() bool {
	return g.Zombie.x == maxX
}

// HitZombie tries to hit the Zombie, it returns boolean
// which describes if the hit was a success.
func (g *Gameboard) HitZombie(x, y int) bool {
	hit := false
	if x == g.Zombie.x && y == g.Zombie.y {
		g.Zombie.Hits++
		hit = true
	}

	return hit
}

// ZombieDead returns whether the Zombie is dead.
func (g *Gameboard) ZombieDead() bool {
	return g.Zombie.Hits >= 3
}
