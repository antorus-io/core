package utils

import (
	"testing"
)

func TestIncludes(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{
			name:     "item exists in slice",
			slice:    []string{"apple", "banana", "cherry"},
			item:     "banana",
			expected: true,
		},
		{
			name:     "item does not exist in slice",
			slice:    []string{"apple", "banana", "cherry"},
			item:     "grape",
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			item:     "apple",
			expected: false,
		},
		{
			name:     "item is first element",
			slice:    []string{"apple", "banana", "cherry"},
			item:     "apple",
			expected: true,
		},
		{
			name:     "item is last element",
			slice:    []string{"apple", "banana", "cherry"},
			item:     "cherry",
			expected: true,
		},
		{
			name:     "slice with duplicate items, item exists",
			slice:    []string{"apple", "banana", "apple"},
			item:     "apple",
			expected: true,
		},
		{
			name:     "slice with duplicate items, item does not exist",
			slice:    []string{"apple", "banana", "apple"},
			item:     "cherry",
			expected: false,
		},
		{
			name:     "empty item",
			slice:    []string{"apple", "", "banana"},
			item:     "",
			expected: true,
		},
		{
			name:     "item is empty, slice is empty",
			slice:    []string{},
			item:     "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Includes(tc.slice, tc.item)
			if actual != tc.expected {
				t.Errorf("Includes(%v, %q) = %v, expected %v", tc.slice, tc.item, actual, tc.expected)
			}
		})
	}
}
