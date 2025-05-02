package database

import (
	"testing"
	"time"

	"github.com/jespino/pool-app/internal/models"
)

func TestMemoryDB(t *testing.T) {
	db := NewMemoryDB()

	// Test creating and retrieving a pool
	pool := models.Pool{
		ID:          "test-pool",
		Title:       "Test Pool",
		Description: "This is a test pool",
		Options: []models.Option{
			{ID: "opt1", Text: "Option 1", Votes: 0},
			{ID: "opt2", Text: "Option 2", Votes: 0},
		},
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// Create the pool
	err := db.CreatePool(pool)
	if err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	// Get the pool
	retrievedPool, err := db.GetPool("test-pool")
	if err != nil {
		t.Fatalf("Failed to get pool: %v", err)
	}

	if retrievedPool.ID != pool.ID {
		t.Errorf("Expected pool ID to be %s, got %s", pool.ID, retrievedPool.ID)
	}

	if retrievedPool.Title != pool.Title {
		t.Errorf("Expected pool title to be %s, got %s", pool.Title, retrievedPool.Title)
	}

	// Test voting
	vote := models.Vote{
		ID:        "vote1",
		PoolID:    "test-pool",
		OptionID:  "opt1",
		UserID:    "user1",
		CreatedAt: time.Now(),
	}

	err = db.CreateVote(vote)
	if err != nil {
		t.Fatalf("Failed to create vote: %v", err)
	}

	// Check if the vote count was incremented
	updatedPool, err := db.GetPool("test-pool")
	if err != nil {
		t.Fatalf("Failed to get updated pool: %v", err)
	}

	if updatedPool.Options[0].Votes != 1 {
		t.Errorf("Expected option 1 to have 1 vote, got %d", updatedPool.Options[0].Votes)
	}

	// Test duplicate vote
	duplicateVote := models.Vote{
		ID:        "vote2",
		PoolID:    "test-pool",
		OptionID:  "opt2",
		UserID:    "user1", // Same user
		CreatedAt: time.Now(),
	}

	err = db.CreateVote(duplicateVote)
	if err == nil {
		t.Errorf("Expected error when user votes twice, but got none")
	}
}