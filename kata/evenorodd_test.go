package kata

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{"Returns Even for 0", 0, "Even"},
		{"Returns Odd for 1", 1, "Odd"},
		{"Returns Even for 2", 2, "Even"},
		{"Returns Odd for -1", -1, "Odd"},
		{"Returns Even for -2", -2, "Even"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := EvenOrOdd(test.arg); got != test.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, test.want)
			}
		})
	}
}
