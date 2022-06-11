package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "15 test 1", input: "a124t000000KDCA", expected: "a124t000000KDCAAA4"},
		{name: "15 test 2", input: "001i000001AWbWu", expected: "001i000001AWbWuAAL"},
		{name: "15 test 3", input: "5003000000D8cuI", expected: "5003000000D8cuIAAR"},
		{name: "18 test", input: "9062I000000XmyrQAC", expected: "9062I000000XmyrQAC"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Convert(tc.input); got != tc.expected {
				t.Errorf("Got %s, expected %s", got, tc.expected)
			}
		})
	}
}
