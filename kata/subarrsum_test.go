package kata

import (
	"fmt"
	"testing"
)

// Returns 6 which is the sum of: [4, -1, 2, 1]
func ExampleSubArrSum_a() {
	sum := SubArrSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(sum)
	// Output: 6
}

// Returns 22 which is the sum of: [2, 9, -4, -3, 8, -10, 20]
func ExampleSubArrSum_b() {
	sum := SubArrSum([]int{5, -6, 2, 9, -4, -3, 8, -10, 20})
	fmt.Println(sum)
	// Output: 22
}

// Returns 36 which is the sum of: [9, 8, 7, -3, 6, 5, 4]
func ExampleSubArrSum_c() {
	sum := SubArrSum([]int{9, 8, 7, -3, 6, 5, 4, -3, 2, 1})
	fmt.Println(sum)
	// Output: 36
}

func TestSubArrSum(t *testing.T) {
	type args struct {
		numbers []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"When list is empty then returns 0", args{[]int{}}, 0},
		{"When list contains only negative numbers then returns 0", args{[]int{-1, -2, -3}}, 0},
		{"When list contains only positive numbers then returns sum of whole list", args{[]int{1, 2, 3}}, 6},
		{"When list contains both positive nad negative numbers then returns max sum of any sub array", args{[]int{-1, -2, -3, 1, 2, 3}}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArrSum(tt.args.numbers); got != tt.want {
				t.Errorf("SubArrSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
