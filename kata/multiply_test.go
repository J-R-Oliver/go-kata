package kata

import (
	"fmt"
	"testing"
)

func ExampleMultiply() {
	product := Multiply(2, 2)
	fmt.Println(product)
	// Output: 4
}

func TestMultiply(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"When passed 2 and 2 then returns 4", args{2, 2}, 4},
		{"When passed 4 and 5 then returns 20", args{4, 5}, 20},
		{"When passed 0 and 100 then returns 0", args{0, 100}, 0},
		{"When passed 1 and 100 then returns 100", args{1, 100}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
