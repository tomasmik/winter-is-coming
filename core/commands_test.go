package core

import (
	"reflect"
	"testing"
)

func TestParseCommandShoot(t *testing.T) {
	type args struct {
		received string
	}
	tests := []struct {
		name    string
		args    args
		want    *CommandShoot
		wantErr bool
	}{
		{
			name: "received wrong command, should error",
			args: args{
				received: "random text",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing parts, should error",
			args: args{
				received: "SHOOT ",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "received non integer x or y, should error",
			args: args{
				received: "SHOOT 0 y",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "received command SHOOT with integer 2 values, should not error",
			args: args{
				received: "SHOOT 2 1",
			},
			want: &CommandShoot{
				X: 2,
				Y: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandShoot(tt.args.received)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandShoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandShoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCommandJoinGame(t *testing.T) {
	type args struct {
		received string
	}
	tests := []struct {
		name    string
		args    args
		want    *CommandJoinGame
		wantErr bool
	}{
		{
			name: "received wrong command, should error",
			args: args{
				received: "random text",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing parts, should error",
			args: args{
				received: "JOINGAME ",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "received command JOINGAME with 1 string arg, should not error",
			args: args{
				received: "JOINGAME mock",
			},
			want: &CommandJoinGame{
				GameName: "mock",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandJoinGame(tt.args.received)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandJoinGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandJoinGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCommandJoinServer(t *testing.T) {
	type args struct {
		received string
	}
	tests := []struct {
		name    string
		args    args
		want    *CommandJoinServer
		wantErr bool
	}{
		{
			name: "received wrong command, should error",
			args: args{
				received: "random text",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing parts, should error",
			args: args{
				received: "JOINSERVER ",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "received command JOINSERVER with 1 string arg, should not error",
			args: args{
				received: "JOINSERVER mock",
			},
			want: &CommandJoinServer{
				Name: "mock",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandJoinServer(tt.args.received)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandJoinServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommandJoinServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCommandType(t *testing.T) {
	type args struct {
		received string
	}
	tests := []struct {
		name    string
		args    args
		want    CommandType
		wantErr bool
	}{
		{
			name: "cant parse command type, should error",
			args: args{
				received: "random text",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "empty string given, should error",
			args: args{
				received: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "command SHOOT, should not error",
			args: args{
				received: "SHOOT 1 1",
			},
			want:    CommandTypeShoot,
			wantErr: false,
		},
		{
			name: "command JOINSERVER, should not error",
			args: args{
				received: "JOINSERVER x",
			},
			want:    CommandTypeJoinServer,
			wantErr: false,
		},
		{
			name: "command JOINGAME, should not error",
			args: args{
				received: "JOINGAME x",
			},
			want:    CommandTypeJoinGame,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommandType(tt.args.received)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommandType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseCommandType() = %v, want %v", got, tt.want)
			}
		})
	}
}
