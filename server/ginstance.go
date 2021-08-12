package server

import (
	"sync"
	"time"

	"github.com/tomasmik/winter-is-coming/core"
)

// gameInstance is a game instance with a player and a zombie
type gameInstance struct {
	name string
	gb   *core.Gameboard

	shotCh chan shot
	respCh chan instanceResp
	// done channel is shared with the keeper.
	// When keeper shuts down, all instances should exit.
	done chan struct{}
	o    sync.Once
}

type shot struct {
	name string
	x    int
	y    int
}

type instanceResp struct {
	IsOver   bool
	GameName string
	Resp     core.Response
}

// Run starts a game instance thread.
// Game instance doesn't know if there are any players
// playing in it, it will run until:
// - A player joins a game instance and wins.
// - Zombie reaches the wall.
// - Application is shutdown.
// Note:
// I got confused regarding this part in the task description
// If 2+ players join, first to shoot wins, but what do we do
// if only one player is in the game and is shooting the zombie
// does it also count at 1 shot wins or do we then allow for more shots
// and change it back to the 2+ player rule when someone joins?
// So this is now a team game against the zombies,
// though I guess it's clearly obviuos that converting to 1 shot wins
// isn't a hard thing to do.
func (g *gameInstance) Run() {
	ticker := time.NewTicker(time.Second * 4)
	defer ticker.Stop()

	// Walk once at the start.
	g.o.Do(func() {
		x, y := g.gb.ZombieWalk()
		g.newMsg(false, core.NewResponseWalk(g.gb.Zombie.Name, x, y))
	})

	for {
		select {
		case <-g.done:
			return
		case shot := <-g.shotCh:
			if g.gb.HitZombie(shot.x, shot.y) {
				g.newMsg(false, core.NewResponseBoom(shot.name, g.gb.Zombie.Name, g.gb.Zombie.Hits))
			}
			if g.gb.ZombieDead() {
				g.newMsg(true, core.NewResponseFinish(true))
				return
			}
		case <-ticker.C:
			x, y := g.gb.ZombieWalk()
			g.newMsg(false, core.NewResponseWalk(g.gb.Zombie.Name, x, y))
			if g.gb.ZombieReachedWall() {
				g.newMsg(true, core.NewResponseFinish(false))
				return
			}
		}
	}
}

func (g *gameInstance) newMsg(isOver bool, resp core.Response) {
	g.respCh <- instanceResp{
		IsOver:   isOver,
		Resp:     resp,
		GameName: g.name,
	}
}

func (g *gameInstance) shoot(name string, x, y int) {
	g.shotCh <- shot{
		name: name,
		x:    x,
		y:    y,
	}
}
