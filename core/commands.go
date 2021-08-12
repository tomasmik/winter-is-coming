package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// CommandJoinServer is returned when a clients
// message is parsed as a request to join a server.
type CommandJoinServer struct {
	Name string
}

// CommandJoinGame is returned when a clients
// message is parsed as a request to join a game.
type CommandJoinGame struct {
	GameName string
}

// CommandShoot is returned when a clients
// message is parsed as a request to shoot an enemy.
type CommandShoot struct {
	X int
	Y int
}

// CommandType is a type which describes the
// possible commands sent by the client to the server.
type CommandType string

const (
	// CommandTypeJoinServer is expected when the client
	// connects and wants to join the server as a player.
	CommandTypeJoinServer CommandType = "JOINSERVER"
	// CommandTypeJoinGame is expected when the client
	// is connected as a plyer and wants to join a game.
	CommandTypeJoinGame CommandType = "JOINGAME"
	// CommandTypeShoot is expected to be received from the client
	// when he has joined a game and is trying to shoot a zombie.
	CommandTypeShoot CommandType = "SHOOT"
)

func ParseCommandShoot(received string) (*CommandShoot, error) {
	parts := strings.Split(received, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("expected format for shoot command is '%s {x} {y}'", CommandTypeShoot)
	}
	if CommandType(parts[0]) != CommandTypeShoot {
		return nil, fmt.Errorf("expected format for shoot command is '%s {x} {y}'", CommandTypeShoot)
	}
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, errors.New("could not parse coordinate x")
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, errors.New("could not parse coordinate y")
	}
	return &CommandShoot{
		X: x,
		Y: y,
	}, nil
}

func ParseCommandJoinGame(received string) (*CommandJoinGame, error) {
	err := fmt.Errorf("expected format for join game command is '%s {name}'", CommandTypeJoinGame)

	parts := strings.Split(received, " ")
	if len(parts) != 2 {
		return nil, err
	}
	if CommandType(parts[0]) != CommandTypeJoinGame {
		return nil, err
	}
	if parts[1] == "" {
		return nil, err
	}
	return &CommandJoinGame{
		GameName: parts[1],
	}, nil
}

func ParseCommandJoinServer(received string) (*CommandJoinServer, error) {
	err := fmt.Errorf("expected format for join command is '%s {name}'", CommandTypeJoinServer)

	parts := strings.Split(received, " ")
	if len(parts) != 2 {
		return nil, err
	}
	if CommandType(parts[0]) != CommandTypeJoinServer {
		return nil, err
	}
	if parts[1] == "" {
		return nil, err
	}
	return &CommandJoinServer{
		Name: parts[1],
	}, nil
}

func ParseCommandType(received string) (CommandType, error) {
	parts := strings.Split(received, " ")
	if len(parts) <= 1 {
		return "", errors.New("a command should consist of type+arguments")
	}

	cmd := CommandType(parts[0])
	switch cmd {
	case CommandTypeShoot, CommandTypeJoinServer, CommandTypeJoinGame:
	default:
		return "", fmt.Errorf("%s is not a command server understands", cmd)
	}
	return cmd, nil
}
