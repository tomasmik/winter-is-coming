package core

import "fmt"

// ResponseBoom is sent back to the client
// if he hits a shot.
type ResponseBoom struct {
	player string
	hits   int
	enemy  string
}

// ResponseWalk is sent to the client
// when a position of an enemy changes.
type ResponseWalk struct {
	enemy string
	x     int
	y     int
}

// ResponseFinish is sent back to the client
// when a game he is in ends.
type ResponseFinish struct {
	won bool
}

// ResponseError is sent to the client when his request
// can't be completed because of an error.
type ResponseError struct {
	err error
}

// ResponseType is a type which describes the
// possible resopnses sent by the server to the client.
type ResponseType string

const (
	// ResponseTypeWalk is expected to be streamed by the server
	// when a zombie moves.
	ResponseTypeWalk ResponseType = "WALK"
	// ResponseTypeBoom is expected to be streamed by the server
	// when a zombie is hit.
	ResponseTypeBoom ResponseType = "BOOM"
	// ResponseError is returned by the server to the client incase
	// of an error.
	ResponseTypeError ResponseType = "ERROR"
	// ResponseFinish is returned by the server to the client
	// when a game ends.
	ResponseTypeFinish ResponseType = "FINISH"
)

// Response interface abstracts away any server
// response, allowing a client to use them without
// having to know any context.
type Response interface {
	// String will convert a response in to a string.
	String() string
}

func NewResponseBoom(player, enemy string, hits int) *ResponseBoom {
	return &ResponseBoom{
		player: player,
		enemy:  enemy,
		hits:   hits,
	}
}

func (r *ResponseBoom) String() string {
	return fmt.Sprintf("%s %s %d %s", ResponseTypeBoom, r.player, r.hits, r.enemy)
}

func NewResponseWalk(enemy string, x, y int) *ResponseWalk {
	return &ResponseWalk{
		enemy: enemy,
		x:     x,
		y:     y,
	}
}

func (r *ResponseWalk) String() string {
	return fmt.Sprintf("%s %s %d %d", ResponseTypeWalk, r.enemy, r.x, r.y)
}

func NewResponseError(err error) *ResponseError {
	return &ResponseError{
		err: err,
	}
}

func (r *ResponseError) String() string {
	return fmt.Sprintf("%s %v", ResponseTypeError, r.err)
}

func NewResponseFinish(won bool) *ResponseFinish {
	return &ResponseFinish{
		won: won,
	}
}

func (r *ResponseFinish) String() string {
	result := "LOST"
	if r.won {
		result = "WON"
	}

	return fmt.Sprintf("%s %s", ResponseTypeFinish, result)
}
