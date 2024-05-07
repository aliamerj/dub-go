package utils

import (
	"net/url"
	"testing"
)

// Define structs to use in testing
type Status string
type Priority int

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"

	PriorityHigh   Priority = 1
	PriorityMedium Priority = 2
	PriorityLow    Priority = 3
)

type TestOptions struct {
	Name     string
	Status   Status
	Priority Priority
	Tags     []string
	page     int
}

// TestBuildQueryURL provides comprehensive testing of various struct inputs
func TestBuildQueryURL(t *testing.T) {
	tests := []struct {
		name     string
		options  interface{}
		expected string
	}{
		{
			name: "Handling enums and custom types",
			options: TestOptions{
				Name:     "Test",
				Status:   StatusActive,
				Priority: PriorityHigh,
				Tags:     []string{"tag1", "tag2"},
				page:     5,
			},
			expected: "name=Test&priority=1&status=active&tags=tag1&tags=tag2",
		},
		{
			name: "Edge cases with enums",
			options: TestOptions{
				Name:     "EdgeCase",
				Status:   StatusInactive,
				Priority: PriorityLow,
				page:     -1,
			},
			expected: "name=EdgeCase&priority=3&status=inactive",
		},
		{
			name: "Empty strings and zero values",
			options: TestOptions{
				Name:     "",
				Status:   "", // Test for empty string as enum
				Priority: 0,  // Test for zero value which might not be a valid enum
				page:     0,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := url.Values{}
			BuildQueryURL(&query, tt.options)

			if query.Encode() != tt.expected {
				t.Errorf("BuildQueryURL() got = %v, want %v", query.Encode(), tt.expected)
			}
		})
	}
}
