package core

import (
	"bitbucket.org/advbet/uid"
)

// Messenger is an object that can be used to
// abtract away client and server communication.
type Messenger struct {
	send      chan Message
	resp      chan Response
	signature uid.UUID
}

// Message is a message sent over the Messenger.
type Message struct {
	DC        bool
	Message   string
	Resp      chan Response
	Signature uid.UUID
}

// NewMessenger returns a new messenger object.
// Given signature must be unique per user.
func NewMessenger(signature uid.UUID, send chan Message) *Messenger {
	return &Messenger{
		signature: signature,
		send:      send,
		resp:      make(chan Response),
	}
}

// SendMessage sends a new message
func (m *Messenger) SendMessage(s string) {
	m.send <- Message{
		Message:   s,
		Resp:      m.resp,
		Signature: m.signature,
	}
}

// Disconnect sends a new disconnet message
// closing the response channel.
// After disconnet is called, messenger is no longer valid.
func (m *Messenger) Disconnect() {
	close(m.resp)
	m.send <- Message{
		DC:        true,
		Resp:      m.resp,
		Signature: m.signature,
	}
}

// Respond sents a response back to the message creator.
func (m *Message) Respond(resp Response) {
	m.Resp <- resp
}

// RespondErr sends a response (error) back to the message creator.
func (m *Message) RespondErr(err error) {
	m.Respond(NewResponseError(err))
}

func (m *Messenger) ReadResponses() <-chan Response {
	return m.resp
}
