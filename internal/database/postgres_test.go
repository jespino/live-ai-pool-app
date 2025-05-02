package database

import (
	"os"
	"testing"
	"time"

	"github.com/jespino/pool-app/internal/models"
)

// This test will be skipped unless POSTGRES_TEST_DSN is set
func TestPostgresDB(t *testing.T) {
	dsn := os.Getenv("POSTGRES_TEST_DSN")
	if dsn == "" {
		t.Skip("Skipping PostgreSQL tests, POSTGRES_TEST_DSN environment variable not set")
	}

	db, err := NewPostgresDB(dsn)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()
	
	// Clean up database before and after test
	if err := db.Cleanup(); err != nil {
		t.Fatalf("Failed to clean up database: %v", err)
	}
	defer db.Cleanup()

	// Initialize the database
	if err := db.Initialize(); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Create test user
	user := models.User{
		ID:       "test-user",
		Username: "testuser",
	}
	if err := db.CreateUser(user); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Create test pool
	pool := models.Pool{
		ID:          "test-pool",
		Title:       "Test Pool",
		Description: "PostgreSQL test pool",
		Options: []models.Option{
			{ID: "opt1", Text: "Option 1", Votes: 0},
			{ID: "opt2", Text: "Option 2", Votes: 0},
		},
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// Test CreatePool
	if err := db.CreatePool(pool); err != nil {
		t.Fatalf("Failed to create pool: %v", err)
	}

	// Test GetPool
	retrievedPool, err := db.GetPool("test-pool")
	if err != nil {
		t.Fatalf("Failed to get pool: %v", err)
	}

	if retrievedPool.ID != pool.ID {
		t.Errorf("Expected pool ID to be %s, got %s", pool.ID, retrievedPool.ID)
	}

	if len(retrievedPool.Options) != 2 {
		t.Errorf("Expected 2 options, got %d", len(retrievedPool.Options))
	}

	// Test CreateVote
	vote := models.Vote{
		ID:        "test-vote",
		PoolID:    "test-pool",
		OptionID:  "opt1",
		UserID:    "test-user",
		CreatedAt: time.Now(),
	}

	if err := db.CreateVote(vote); err != nil {
		t.Fatalf("Failed to create vote: %v", err)
	}

	// Check if vote was counted
	updatedPool, err := db.GetPool("test-pool")
	if err != nil {
		t.Fatalf("Failed to get updated pool: %v", err)
	}

	var opt1Votes int
	for _, opt := range updatedPool.Options {
		if opt.ID == "opt1" {
			opt1Votes = opt.Votes
			break
		}
	}

	if opt1Votes != 1 {
		t.Errorf("Expected option 1 to have 1 vote, got %d", opt1Votes)
	}

	// Test getting votes for pool
	votes, err := db.GetVotesForPool("test-pool")
	if err != nil {
		t.Fatalf("Failed to get votes: %v", err)
	}

	if len(votes) != 1 {
		t.Errorf("Expected 1 vote, got %d", len(votes))
	}

	// Test DeletePool
	if err := db.DeletePool("test-pool"); err != nil {
		t.Fatalf("Failed to delete pool: %v", err)
	}

	// Verify pool is deleted
	_, err = db.GetPool("test-pool")
	if err == nil {
		t.Errorf("Expected error when getting deleted pool, got nil")
	}
}