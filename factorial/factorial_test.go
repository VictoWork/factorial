package factorial

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},        // Factorial of 0 is 1
		{1, 1},        // Factorial of 1 is 1
		{2, 2},        // Factorial of 2 is 2
		{3, 6},        // Factorial of 3 is 6
		{4, 24},       // Factorial of 4 is 24
		{5, 120},      // Factorial of 5 is 120
		{6, 720},      // Factorial of 6 is 720
		{10, 3628800}, // Factorial of 10 is 3628800
	}

	for _, test := range tests {
		if output := Factorial(test.input); output != test.expected {
			t.Errorf("Factorial(%d) = %d; want %d", test.input, output, test.expected)
		}
	}
}
