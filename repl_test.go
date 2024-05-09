package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
        expected []string
	} {
		{
			input: " Hello   ",
			expected: []string{"hello"},
		}, 
		{
			input: "    ",
			expected: []string{},
		},
		{
			input: "    hello   world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "helLo world",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := getCleanInput(c.input)
        if len(actual)!= len(c.expected)  {
			t.Errorf("lenghts don't match expected: %v, got %v", len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord!= expectedWord {
                t.Errorf("words don't match expected: %v, got %v", expectedWord, actualWord)
            }
		}
	}
}