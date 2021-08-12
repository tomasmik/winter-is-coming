package core

import (
	"errors"
	"fmt"
	"testing"
)

func TestResponseBoom_String(t *testing.T) {
	mockStr1 := "A"
	mockStr2 := "B"
	mockInt1 := 1
	type fields struct {
		player string
		hits   int
		enemy  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "to string ResponseBoom",
			fields: fields{
				player: mockStr1,
				hits:   mockInt1,
				enemy:  mockStr2,
			},
			want: fmt.Sprintf("%s %s %d %s", ResponseTypeBoom, mockStr1, mockInt1, mockStr2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResponseBoom{
				player: tt.fields.player,
				hits:   tt.fields.hits,
				enemy:  tt.fields.enemy,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ResponseBoom.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseWalk_String(t *testing.T) {
	mockStr1 := "A"
	mockInt1 := 1
	mockInt2 := 2
	type fields struct {
		enemy string
		x     int
		y     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "to string ResponseWalk",
			fields: fields{
				enemy: mockStr1,
				x:     mockInt1,
				y:     mockInt2,
			},
			want: fmt.Sprintf("%s %s %d %d", ResponseTypeWalk, mockStr1, mockInt1, mockInt2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResponseWalk{
				enemy: tt.fields.enemy,
				x:     tt.fields.x,
				y:     tt.fields.y,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ResponseWalk.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseError_String(t *testing.T) {
	mockErr := errors.New("a")
	type fields struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "to string ResponseWalk",
			fields: fields{
				err: mockErr,
			},
			want: fmt.Sprintf("%s %v", ResponseTypeError, mockErr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResponseError{
				err: tt.fields.err,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ResponseError.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseFinish_String(t *testing.T) {
	mockBoolt := true
	mockBoolf := false
	type fields struct {
		won bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "to string ResponseFinish won",
			fields: fields{
				won: mockBoolt,
			},
			want: fmt.Sprintf("%s %s", ResponseTypeFinish, "WON"),
		},
		{
			name: "to string ResponseFinish lost",
			fields: fields{
				won: mockBoolf,
			},
			want: fmt.Sprintf("%s %s", ResponseTypeFinish, "LOST"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResponseFinish{
				won: tt.fields.won,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("ResponseFinish.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
