package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  HELLO  ",
			expected: []string{"hello", "hello"},
		},

		{
			input:    "  hello  HELLO tresCuATRO  ",
			expected: []string{"hello", "hello", "trescuatro"},
		},

		{
			input:    "",
			expected: []string{},
		},
		// add more cases here
		{
			input:    "  hello  HELLO hay ",
			expected: []string{"hello", "hello", "hay"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("longitudes no concuerdan: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) -> %v. se esperaba: %v", c.input, actual, c.expected)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}

}
