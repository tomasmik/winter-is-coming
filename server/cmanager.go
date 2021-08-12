package server

import (
	"net"
	"sync"
)

type Cmanager struct {
	// A map might be a little bit less efficient
	// but it allows us to not keep track of the
	// indexes at which clients are placed, so removing
	// them is much easier.
	clients map[net.Conn]struct{}
	m       sync.Mutex
}

func NewCmanager() *Cmanager {
	return &Cmanager{
		clients: make(map[net.Conn]struct{}),
	}
}

func (m *Cmanager) Reg(c net.Conn) {
	m.m.Lock()
	defer m.m.Unlock()
	m.clients[c] = struct{}{}
}

func (m *Cmanager) Unreg(c net.Conn) {
	m.m.Lock()
	defer m.m.Unlock()

	c.Close()
	delete(m.clients, c)
}

func (m *Cmanager) UnregAll() {
	m.m.Lock()
	defer m.m.Unlock()

	for c := range m.clients {
		c.Close()
		delete(m.clients, c)
	}
}
