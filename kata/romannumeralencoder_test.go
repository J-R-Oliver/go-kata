package kata

import (
	"fmt"
	"testing"
)

// Returns III which is the Roman Numeral of 3
func ExampleRomanNumeralEncoder_3() {
	r := RomanNumeralEncoder(3)
	fmt.Println(r)
	// Output: III
}

// Returns XLII which is the Roman Numeral of 42
func ExampleRomanNumeralEncoder_42() {
	r := RomanNumeralEncoder(42)
	fmt.Println(r)
	// Output: XLII
}

// Returns XCVIII which is the Roman Numeral of 98
func ExampleRomanNumeralEncoder_98() {
	r := RomanNumeralEncoder(98)
	fmt.Println(r)
	// Output: XCVIII
}

func TestRomanNumeralEncoder(t *testing.T) {
	type args struct {
		number int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"When passed 1 then return is I", args{1}, "I"},
		{"When passed 2 then return is II", args{2}, "II"},
		{"When passed 4 then return is IV", args{4}, "IV"},
		{"When passed 5 then return is V", args{5}, "V"},
		{"When passed 9 then return is IX", args{9}, "IX"},
		{"When passed 10 then return is X", args{10}, "X"},
		{"When passed 13 then return is XIII", args{13}, "XIII"},
		{"When passed 17 then return is XVII", args{17}, "XVII"},
		{"When passed 19 then return is XIX", args{19}, "XIX"},
		{"When passed 39 then return is XXXIX ", args{39}, "XXXIX"},
		{"When passed 40 then return is XL ", args{40}, "XL"},
		{"When passed 50 then return is L ", args{50}, "L"},
		{"When passed 60 then return is LX", args{60}, "LX"},
		{"When passed 75 then return is LXXV", args{75}, "LXXV"},
		{"When passed 91 then return is XCI", args{91}, "XCI"},
		{"When passed 100 then return is C", args{100}, "C"},
		{"When passed 198 then return is CXCVIII", args{198}, "CXCVIII"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanNumeralEncoder(tt.args.number); got != tt.want {
				t.Errorf("RomanNumeralEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
