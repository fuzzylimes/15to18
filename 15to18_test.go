package sf15to18

import (
	"testing"
)

func TestGetArgs(t *testing.T) {
	testCases := []struct {
		name     string
		filepath string
		cliArgs  []string
		expected []string
	}{
		{name: "CLI args only", filepath: "", cliArgs: []string{"a124t000000KDCA", "001i000001AWbWu"}, expected: []string{"a124t000000KDCA", "001i000001AWbWu"}},
		{name: "filepath only", filepath: "./test/sample.json", cliArgs: nil, expected: []string{"{\n  \"environment\": [\n    {\n      \"name\": \"dev\",\n      \"id\": \"a124t000000KDCA\"\n    },\n    {\n      \"name\": \"qa\",\n      \"id\": \"001i000001AWbWu\"\n    },\n    {\n      \"name\": \"prod\",\n      \"id\": \"5003000000D8cuI\"\n    }\n  ]\n}"}},
		{name: "filepath and CLI args", filepath: "./test/sample.json", cliArgs: []string{"a124t000000KDCA", "001i000001AWbWu"}, expected: []string{"{\n  \"environment\": [\n    {\n      \"name\": \"dev\",\n      \"id\": \"a124t000000KDCA\"\n    },\n    {\n      \"name\": \"qa\",\n      \"id\": \"001i000001AWbWu\"\n    },\n    {\n      \"name\": \"prod\",\n      \"id\": \"5003000000D8cuI\"\n    }\n  ]\n}"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := HandleArgs(tc.filepath, tc.cliArgs); !equal(got, tc.expected) {
				t.Errorf("Got %q, expected %q", got, tc.expected)
			}
		})
	}
}

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

func TestValidIds(t *testing.T) {
	testCases := []struct {
		name     string
		in       []string
		expected []string
	}{
		{name: "all valid", in: []string{"a124t000000KDCA", "001i000001AWbWu", "5003000000D8cuI"}, expected: []string{"a124t000000KDCA", "001i000001AWbWu", "5003000000D8cuI"}},
		{name: "mixed", in: []string{"a124t0_0000KDCA", "001i000001AWbWu", "5003000000D8cuIAAD"}, expected: []string{"001i000001AWbWu", "5003000000D8cuI"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := GetValidIds(tc.in); !equal(got, tc.expected) {
				t.Errorf("Got %q, expected %q", got, tc.expected)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
