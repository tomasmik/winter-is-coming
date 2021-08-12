package server

import (
	"bufio"
	"errors"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"bitbucket.org/advbet/uid"
	"github.com/sirupsen/logrus"
	"github.com/tomasmik/winter-is-coming/core"
)

// Server can be used to manage connections.
// It wraps the listener allowing it to accept new connection
// and keeps track of connected clients.
type Server struct {
	l        net.Listener
	cmanager *Cmanager
	gp       *GameKeeper

	done chan struct{}
	cwg  sync.WaitGroup
	log  *logrus.Entry
}

// New creates a new tcp connection and returns
// a new server object which can be used to manager that connection.
func New(l net.Listener) *Server {
	return &Server{
		l:        l,
		gp:       NewGameKeeper(),
		cmanager: NewCmanager(),
		done:     make(chan struct{}, 0),
		log:      logrus.WithField("thread", "tcp-server"),
	}
}

// Run the server. Allowing it to accept new connections.
// The server is stopped when the Stop() method is called.
func (s *Server) Run() {
	s.log.Info("started")
	defer s.log.Info("stopped")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.gp.Run()
	}()

	for {
		select {
		case <-s.done:
			// Disconnets all users.
			s.cmanager.UnregAll()
			s.cwg.Wait()

			// Stop all games
			s.gp.Stop()
			wg.Wait()
			return
		default:
			c, err := s.l.Accept()
			if err != nil {
				// Dont log an error if the connection is closed because
				// and we're stopping as we've killed the listener.
				if !s.stopped() {
					s.log.WithError(err).Error("accepting new clients")
				}
				break
			}

			s.cmanager.Reg(c)
			s.cwg.Add(1)
			go func(c net.Conn) {
				defer func() {
					// If client is stopped, it'll clean up itself.
					if !s.stopped() {
						s.cmanager.Unreg(c)
					}
				}()

				defer s.cwg.Done()
				s.handleConnection(c)
			}(c)
		}
	}
}

// listen will listen for any incoming messages and pass them along the send channel.
// This func should block until the conection is closed or thread is stopped.
func (s *Server) listen(c net.Conn, p *core.Messenger) {
	for {
		c.SetReadDeadline(time.Now().Add(time.Second * 60))
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) && !s.stopped() {
				s.log.WithError(err).Error("reading from a connection")
			}
			return
		}
		p.SendMessage(strings.TrimSpace(string(msg)))
	}
}

// write writes the received information to the connection.
func (s *Server) write(c net.Conn, p *core.Messenger, stopped chan struct{}) {
	// This is kinda hacky, but i guess ok for such a thing.
	defer close(stopped)

	for msg := range p.ReadResponses() {
		c.SetWriteDeadline(time.Now().Add(time.Second * 5))
		if _, err := c.Write([]byte(msg.String() + "\n")); err != nil {
			s.log.WithError(err).Error("writing to a connection")
		}
	}
}

func (s *Server) handleConnection(c net.Conn) {
	p := s.gp.NewConnection(uid.NewTimeRand())

	stopped := make(chan struct{})
	go s.write(c, p, stopped)
	s.listen(c, p)

	p.Disconnect()
	<-stopped
}

// stopped can be used to figure out if the thread has stopped.
func (s *Server) stopped() bool {
	select {
	case <-s.done:
		return true
	default:
		return false
	}
}

// Stop will stop the server thread, closing the
// net.Listener and all the clients connections.
func (s *Server) Stop() {
	close(s.done)
	s.l.Close()
}
