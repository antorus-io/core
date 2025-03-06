package utils

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	// Test cases
	testCases := []struct {
		name      string
		input     []int
		transform func(int) string
		expected  []string
	}{
		{
			name:      "square to string",
			input:     []int{1, 2, 3, 4, 5},
			transform: func(i int) string { return strconv.Itoa(i * i) },
			expected:  []string{"1", "4", "9", "16", "25"},
		},
		{
			name:      "empty slice",
			input:     []int{},
			transform: func(i int) string { return strconv.Itoa(i) },
			expected:  []string{},
		},
		{
			name:      "negative numbers",
			input:     []int{-1, -2, -3},
			transform: func(i int) string { return strconv.Itoa(-i) },
			expected:  []string{"1", "2", "3"},
		},
		{
			name:      "zero",
			input:     []int{0},
			transform: func(i int) string { return strconv.Itoa(i) },
			expected:  []string{"0"},
		},
		{
			name:      "Mixed numbers",
			input:     []int{-2, 0, 2},
			transform: func(i int) string { return strconv.Itoa(i) },
			expected:  []string{"-2", "0", "2"},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Map(tc.input, tc.transform)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Map(%v, transform) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}

	//Test with strings
	stringTestCases := []struct {
		name      string
		input     []string
		transform func(string) int
		expected  []int
	}{
		{
			name:      "string length",
			input:     []string{"a", "ab", "abc"},
			transform: func(s string) int { return len(s) },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "empty string slice",
			input:     []string{},
			transform: func(s string) int { return len(s) },
			expected:  []int{},
		},
	}

	for _, tc := range stringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Map(tc.input, tc.transform)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Map(%v, transform) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
