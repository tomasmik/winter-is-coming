package core

import (
	"testing"
)

func TestGameboard_ZombieWalk(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "zombie should walk +1 either x or y",
			fields: fields{
				x: 0,
				y: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gameboard{
				Zombie: &Zombie{
					x: 0,
					y: 0,
				},
			}
			// Sort of annoying to test considering that we have
			// a random seed.
			got, got1 := g.ZombieWalk()
			if got == tt.fields.x && got1 == tt.fields.y {
				t.Errorf("Gameboard.ZombieWalk() didn't update neither x or y")
			}
		})
	}
}

func TestGameboard_ZombieReachedWall(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zombie has reached the wall, should return true",
			fields: fields{
				x: maxX,
				y: 0,
			},
			want: true,
		},
		{
			name: "zombie has not reached the wall, should return false",
			fields: fields{
				x: 0,
				y: 0,
			},
			want: false,
		},
		{
			name: "zombie has not reached the wall, should return false",
			fields: fields{
				x: 0,
				y: maxY,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gameboard{
				Zombie: &Zombie{
					x: tt.fields.x,
					y: tt.fields.y,
				},
			}
			if got := g.ZombieReachedWall(); got != tt.want {
				t.Errorf("Gameboard.ZombieReachedWall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameboard_ZombieHit(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "hit was a success should return true",
			fields: fields{
				x: 0,
				y: 1,
			},
			args: args{
				x: 0,
				y: 1,
			},
			want: true,
		},
		{
			name: "hit was a falure should return false",
			fields: fields{
				x: 1,
				y: 1,
			},
			args: args{
				x: 0,
				y: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gameboard{
				Zombie: &Zombie{
					x: tt.fields.x,
					y: tt.fields.y,
				},
			}
			hit := g.HitZombie(tt.args.x, tt.args.y)
			if hit != tt.want {
				t.Errorf("Gameboard.HitZombie() = %v, want %v", hit, tt.want)
			}
			if hit && g.Zombie.Hits == 0 {
				t.Errorf("Gameboard.HitZombie() should increments hits")
			}
		})
	}
}

func TestGameboard_ZombieDead(t *testing.T) {
	type fields struct {
		Hits int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zombie isn't dead yet, should return false",
			fields: fields{
				Hits: 1,
			},
			want: false,
		},
		{
			name: "zombie is dead, should return true",
			fields: fields{
				Hits: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gameboard{
				Zombie: &Zombie{
					Hits: tt.fields.Hits,
				},
			}
			if got := g.ZombieDead(); got != tt.want {
				t.Errorf("Gameboard.ZombieDead() = %v, want %v", got, tt.want)
			}
		})
	}
}
