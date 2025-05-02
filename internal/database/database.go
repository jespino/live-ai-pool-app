package database

import (
	"github.com/jespino/pool-app/internal/models"
)

// Storage is the interface that all database implementations must satisfy
type Storage interface {
	// Pool operations
	CreatePool(pool models.Pool) error
	GetPool(id string) (models.Pool, error)
	ListPools() []models.Pool
	UpdatePool(pool models.Pool) error
	DeletePool(id string) error

	// Vote operations
	CreateVote(vote models.Vote) error
	GetVotesForPool(poolID string) ([]models.Vote, error)

	// User operations
	CreateUser(user models.User) error
	GetUser(id string) (models.User, error)
}