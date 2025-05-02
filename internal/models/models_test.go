package models

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// Create a test pool
	options := []Option{
		{ID: "opt1", Text: "Option 1", Votes: 0},
		{ID: "opt2", Text: "Option 2", Votes: 0},
	}
	
	now := time.Now()
	expiry := now.Add(24 * time.Hour)
	
	pool := Pool{
		ID:          "test-pool",
		Title:       "Test Pool",
		Description: "This is a test pool",
		Options:     options,
		CreatedAt:   now,
		ExpiresAt:   expiry,
	}
	
	// Check the pool properties
	if pool.ID != "test-pool" {
		t.Errorf("Expected pool ID to be 'test-pool', got %s", pool.ID)
	}
	
	if pool.Title != "Test Pool" {
		t.Errorf("Expected pool title to be 'Test Pool', got %s", pool.Title)
	}
	
	if len(pool.Options) != 2 {
		t.Errorf("Expected pool to have 2 options, got %d", len(pool.Options))
	}
	
	// Check option properties
	if pool.Options[0].ID != "opt1" || pool.Options[0].Text != "Option 1" {
		t.Errorf("First option does not match expected values")
	}
	
	if pool.Options[1].ID != "opt2" || pool.Options[1].Text != "Option 2" {
		t.Errorf("Second option does not match expected values")
	}
}