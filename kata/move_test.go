package kata

import "testing"

func TestMove(t *testing.T) {
	type args struct {
		position int
		roll     int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Returns 8 for a starting position of 0 and a roll of 4", args{0, 4}, 8},
		{"Returns 15 for a starting position of 3 and a roll of 6", args{3, 6}, 15},
		{"Returns 12 for a starting position of 2 and a roll of 5", args{2, 5}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Move(tt.args.position, tt.args.roll); got != tt.want {
				t.Errorf("Move() = %v, want %v", got, tt.want)
			}
		})
	}
}
