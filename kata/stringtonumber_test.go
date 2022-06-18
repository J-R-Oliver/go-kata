package kata

import (
	"fmt"
	"testing"
)

// Returns integer 42 when passed a string of "42"
func ExampleStringToNumber() {
	n := StringToNumber("42")
	fmt.Println(n)
	// Output: 42
}

func TestStringToNumber(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{`When passed "1" then returns 1`, args{"1"}, 1},
		{`When passed "42" then returns 42`, args{"42"}, 42},
		{`When passed "569" then returns 569`, args{"569"}, 569},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToNumber(tt.args.str); got != tt.want {
				t.Errorf("StringToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
