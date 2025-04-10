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
		}, {
			input:    "I am Batman!!",
			expected: []string{"I", "am", "batman!!"},
		}, {
			input:    "Say there, how do YOU do? I... Am not sure!!",
			expected: []string{"say", "there,", "how", "do", "you", "do?", "I...", "am", "not", "sure!!"},
		},
	}

	// for _, c := range cases {
	// 	actual := cleanInput(c.input)
	//   if

	// 	for i := range actual {
	// 		word := actual[i]
	// 		expectedWord := c.expected[i]
	// 		// Check each word in the slice
	// 		// if they don't match, use t.Errorf to print an error message
	// 		// and fail the test
	// 	}
	// }
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(`Test Failed. Expected: %v, Received: %v`, c.expected, actual)
			t.Errorf(`Expected Slice length: %v, Received: %v`, len(actual), len(c.expected))
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test failed, Expected: %s, Received: %s", expectedWord, word)
			}
		}
	}
}
