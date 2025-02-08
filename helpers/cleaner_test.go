package helper

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " CHARmander bulbasaur squirtle",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned slice of length %d, expected %d", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("cleanInput(%q) at index %d returned %q, expected %q",
					c.input, i, actualWord, expectedWord)
			}

		}
	}

}
