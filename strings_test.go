package strings

import (
	"errors"
	"strings"
	"testing"
)

// TestFlip unit test Flip function.
func TestFlip(t *testing.T) {
	testCases := []struct {
		input, expected string
	}{
		{"GO je úžasné!", "!énsažú ej OG"},
		{"golang", "gnalog"},
		{"", ""},
		{"x", "x"},
	}
	for _, test := range testCases {
		observed := Flip(test.input)
		if strings.Compare(test.expected, observed) != 0 {
			t.Errorf("for input '%s', expected '%s', observed '%s'\n", test.input, test.expected, observed)
		}
	}
}

// BenchmarkFlip benchmark Flip function.
func BenchmarkFlip(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_ = Flip("GO je úžasné!")
	}
}

// TestPigLatin unit test PigLatin function.
func TestPigLatin(t *testing.T) {
	testCases := []struct {
		input, expected string
		err             error
	}{
		{"pig", "igpay", nil},
		{"banana", "ananabay", nil},
		{"duck", "uckday", nil},
		{"adi", "adiway", nil},
		{"", "", errors.New("cannot translate an empty string")},
	}
	for _, test := range testCases {
		observed, err := PigLatin(test.input)
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("for input '%s', expected error '%s', observed error '%s'",
					test.input, test.err.Error(), err.Error())
			}
		}
		if observed != test.expected {
			t.Errorf("for input '%s', expected '%s', observed '%s'",
				test.input, test.expected, observed)
		}
	}
}

// BenchmarkPigLation benchmark PigLation function.
func BenchmarkPigLatin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PigLatin("duck")
	}
}
