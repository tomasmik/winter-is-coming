package server

import (
	"errors"
	"sync"

	"bitbucket.org/advbet/uid"
	"github.com/sirupsen/logrus"
	"github.com/tomasmik/winter-is-coming/core"
)

// GameKeeper is used to manage game instances.
// It tracks players and instances that they are in.
type GameKeeper struct {
	players   map[uid.UUID]core.Player
	instances map[string]gameInstance

	umsg chan core.Message
	gmsg chan instanceResp

	done chan struct{}
	log  *logrus.Entry
	iwg  sync.WaitGroup
}

var (
	errNoSession   = errors.New("haven't created a session")
	errHaveSession = errors.New("already created a session")
	errNotInGame   = errors.New("not in a game")
	errNameTaken   = errors.New("name taken")
)

// NewGameKeeper returns a GameKeeper object.
func NewGameKeeper() *GameKeeper {
	return &GameKeeper{
		players:   make(map[uid.UUID]core.Player),
		instances: make(map[string]gameInstance),
		gmsg:      make(chan instanceResp, 16),
		umsg:      make(chan core.Message, 16),
		log:       logrus.WithField("thread", "game-keeper"),
		done:      make(chan struct{}),
	}
}

// NewConnection allows the method user to establish a new connection
// to the game keeper. It returns a messenger object
// using which the user can interact with the game keeper.
//
// Provided argument `sign` must be a unique user session identifier.
func (g *GameKeeper) NewConnection(sign uid.UUID) *core.Messenger {
	return core.NewMessenger(sign, g.umsg)
}

// Run starts a process that manages game instances and players.
func (g *GameKeeper) Run() {
	g.log.Info("started")
	defer g.log.Info("stopped")

	cleanup := func(sign uid.UUID) {
		_, ok := g.players[sign]
		if !ok {
			return
		}
		delete(g.players, sign)
	}

	for {
		select {
		case <-g.done:
			g.iwg.Wait()
			return
		case msg := <-g.gmsg:
			// TODO: This is not really efficient and I am
			// aware of this, but this is the easiest way
			// and it works OK while we dont have a million users.
			for sign, p := range g.players {
				if p.GameName != msg.GameName {
					continue
				}

				p.Resp <- msg.Resp
				if msg.IsOver {
					p.GameName = ""
					g.players[sign] = p
				}
			}
			if msg.IsOver {
				delete(g.instances, msg.GameName)
			}
		case msg := <-g.umsg:
			if msg.DC {
				cleanup(msg.Signature)
				break
			}

			typ, err := core.ParseCommandType(msg.Message)
			if err != nil {
				msg.Respond(core.NewResponseError(err))
			}

			switch typ {
			case core.CommandTypeJoinServer:
				g.msgJoinServer(&msg)
			case core.CommandTypeJoinGame:
				g.msgJoinGame(&msg)
			case core.CommandTypeShoot:
				g.msgShoot(&msg)
			}
		}
	}
}

func (g *GameKeeper) newGameInstance(name string) *gameInstance {
	return &gameInstance{
		name:   name,
		gb:     core.NewGameBoard(),
		shotCh: make(chan shot),
		respCh: g.gmsg,
		done:   g.done,
	}
}

func (g *GameKeeper) msgJoinServer(msg *core.Message) {
	if _, ok := g.players[msg.Signature]; ok {
		msg.RespondErr(errHaveSession)
		return
	}

	cmd, err := core.ParseCommandJoinServer(msg.Message)
	if err != nil {
		msg.RespondErr(err)
		return
	}

	// TODO: I guess it'd be nice to not do it this way.
	for _, p := range g.players {
		if p.Name != cmd.Name {
			continue
		}
		msg.RespondErr(errNameTaken)
		return
	}

	player := core.NewPlayer(cmd.Name, msg.Signature, msg.Resp)
	g.players[msg.Signature] = *player
}

func (g *GameKeeper) msgJoinGame(msg *core.Message) {
	cmd, err := core.ParseCommandJoinGame(msg.Message)
	if err != nil {
		msg.RespondErr(err)
		return
	}

	p, ok := g.players[msg.Signature]
	if !ok {
		msg.RespondErr(errNoSession)
		return
	}

	p.GameName = cmd.GameName
	g.players[msg.Signature] = p

	// If a game instance with this name already exists
	// dont start a new thread.
	if _, ok := g.instances[cmd.GameName]; ok {
		return
	}
	gin := g.newGameInstance(cmd.GameName)
	g.instances[gin.name] = *gin

	g.iwg.Add(1)
	go func() {
		defer g.iwg.Done()
		gin.Run()
	}()
}

func (g *GameKeeper) msgShoot(msg *core.Message) {
	p, ok := g.players[msg.Signature]
	if !ok {
		msg.RespondErr(errNoSession)
		return
	}

	if p.GameName == "" {
		msg.RespondErr(errNotInGame)
		return
	}

	cmd, err := core.ParseCommandShoot(msg.Message)
	if err != nil {
		msg.RespondErr(err)
		return
	}
	gin := g.instances[p.GameName]
	gin.shoot(p.Name, cmd.X, cmd.Y)
}

// Stop will stop the game streamer
func (g *GameKeeper) Stop() {
	close(g.done)
}
