package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "I am Batman!!",
			expected: []string{"i", "am", "batman!!"},
		}, {
			input:    "Say there, how do YOU do? I... Am not sure!!",
			expected: []string{"say", "there,", "how", "do", "you", "do?", "i...", "am", "not", "sure!!"},
		},
	}

	for _, c := range cases {
		received := cleanInput(c.input)
		fmt.Printf("Expected: %q\n", c.expected)
		fmt.Printf("Received: %q\n", received)
		if len(received) != len(c.expected) {
			t.Errorf(`Test Failed. Expected: %q, Received: %q`, c.expected, received)
			t.Errorf(`Expected Slice length: %q, Received: %q`, len(received), len(c.expected))
			t.Fail()
		}
		for i := range received {
			word := received[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test failed, Expected: %s, Received: %s", expectedWord, word)
				t.Fail()
			}
		}
	}
}
