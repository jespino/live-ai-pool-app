package database

import (
	"errors"
	"sync"

	"github.com/jespino/pool-app/internal/models"
)

type MemoryDB struct {
	pools    map[string]models.Pool
	votes    map[string]models.Vote
	users    map[string]models.User
	poolLock sync.RWMutex
	voteLock sync.RWMutex
	userLock sync.RWMutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		pools: make(map[string]models.Pool),
		votes: make(map[string]models.Vote),
		users: make(map[string]models.User),
	}
}

func (db *MemoryDB) CreatePool(pool models.Pool) error {
	db.poolLock.Lock()
	defer db.poolLock.Unlock()

	if _, exists := db.pools[pool.ID]; exists {
		return errors.New("pool already exists")
	}

	db.pools[pool.ID] = pool
	return nil
}

func (db *MemoryDB) GetPool(id string) (models.Pool, error) {
	db.poolLock.RLock()
	defer db.poolLock.RUnlock()

	pool, exists := db.pools[id]
	if !exists {
		return models.Pool{}, errors.New("pool not found")
	}

	return pool, nil
}

func (db *MemoryDB) ListPools() []models.Pool {
	db.poolLock.RLock()
	defer db.poolLock.RUnlock()

	pools := make([]models.Pool, 0, len(db.pools))
	for _, pool := range db.pools {
		pools = append(pools, pool)
	}

	return pools
}

func (db *MemoryDB) UpdatePool(pool models.Pool) error {
	db.poolLock.Lock()
	defer db.poolLock.Unlock()

	if _, exists := db.pools[pool.ID]; !exists {
		return errors.New("pool not found")
	}

	db.pools[pool.ID] = pool
	return nil
}

func (db *MemoryDB) DeletePool(id string) error {
	db.poolLock.Lock()
	defer db.poolLock.Unlock()

	if _, exists := db.pools[id]; !exists {
		return errors.New("pool not found")
	}

	delete(db.pools, id)
	return nil
}

func (db *MemoryDB) CreateVote(vote models.Vote) error {
	db.voteLock.Lock()
	defer db.voteLock.Unlock()
	db.poolLock.Lock()
	defer db.poolLock.Unlock()

	if _, exists := db.votes[vote.ID]; exists {
		return errors.New("vote already exists")
	}

	pool, exists := db.pools[vote.PoolID]
	if !exists {
		return errors.New("pool not found")
	}

	// Check if user already voted for this pool
	for _, v := range db.votes {
		if v.PoolID == vote.PoolID && v.UserID == vote.UserID {
			return errors.New("user already voted in this pool")
		}
	}

	// Find the option and increment its vote count
	optionFound := false
	for i, option := range pool.Options {
		if option.ID == vote.OptionID {
			pool.Options[i].Votes++
			optionFound = true
			break
		}
	}

	if !optionFound {
		return errors.New("option not found in pool")
	}

	db.votes[vote.ID] = vote
	db.pools[vote.PoolID] = pool

	return nil
}

func (db *MemoryDB) GetVotesForPool(poolID string) ([]models.Vote, error) {
	db.voteLock.RLock()
	defer db.voteLock.RUnlock()

	var poolVotes []models.Vote
	for _, vote := range db.votes {
		if vote.PoolID == poolID {
			poolVotes = append(poolVotes, vote)
		}
	}

	return poolVotes, nil
}

func (db *MemoryDB) CreateUser(user models.User) error {
	db.userLock.Lock()
	defer db.userLock.Unlock()

	if _, exists := db.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	db.users[user.ID] = user
	return nil
}

func (db *MemoryDB) GetUser(id string) (models.User, error) {
	db.userLock.RLock()
	defer db.userLock.RUnlock()

	user, exists := db.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}