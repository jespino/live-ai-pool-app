package database

import (
	"database/sql"
	"errors"

	"github.com/jespino/pool-app/internal/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(connectionString string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) Initialize() error {
	// Create pools table
	_, err := p.db.Exec(`
	CREATE TABLE IF NOT EXISTS pools (
		id VARCHAR(36) PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP NOT NULL,
		expires_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		return err
	}

	// Create options table
	_, err = p.db.Exec(`
	CREATE TABLE IF NOT EXISTS options (
		id VARCHAR(36) PRIMARY KEY,
		pool_id VARCHAR(36) NOT NULL REFERENCES pools(id) ON DELETE CASCADE,
		text TEXT NOT NULL,
		votes INT DEFAULT 0,
		UNIQUE(pool_id, id)
	)`)
	if err != nil {
		return err
	}

	// Create users table
	_, err = p.db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36) PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE
	)`)
	if err != nil {
		return err
	}

	// Create votes table
	_, err = p.db.Exec(`
	CREATE TABLE IF NOT EXISTS votes (
		id VARCHAR(36) PRIMARY KEY,
		pool_id VARCHAR(36) NOT NULL REFERENCES pools(id) ON DELETE CASCADE,
		option_id VARCHAR(36) NOT NULL REFERENCES options(id) ON DELETE CASCADE,
		user_id VARCHAR(36) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		created_at TIMESTAMP NOT NULL,
		UNIQUE(pool_id, user_id)
	)`)
	return err
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}

// Cleanup drops all tables - useful for testing
func (p *PostgresDB) Cleanup() error {
	_, err := p.db.Exec(`
	DROP TABLE IF EXISTS votes;
	DROP TABLE IF EXISTS options;
	DROP TABLE IF EXISTS pools;
	DROP TABLE IF EXISTS users;
	`)
	return err
}

func (p *PostgresDB) CreatePool(pool models.Pool) error {
	// Start a transaction
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert into pools table
	_, err = tx.Exec(
		"INSERT INTO pools (id, title, description, created_at, expires_at) VALUES ($1, $2, $3, $4, $5)",
		pool.ID, pool.Title, pool.Description, pool.CreatedAt, pool.ExpiresAt,
	)
	if err != nil {
		return err
	}

	// Insert options
	for _, option := range pool.Options {
		_, err = tx.Exec(
			"INSERT INTO options (id, pool_id, text, votes) VALUES ($1, $2, $3, $4)",
			option.ID, pool.ID, option.Text, option.Votes,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (p *PostgresDB) GetPool(id string) (models.Pool, error) {
	// Query the pool
	pool := models.Pool{}
	err := p.db.QueryRow(
		"SELECT id, title, description, created_at, expires_at FROM pools WHERE id = $1",
		id,
	).Scan(&pool.ID, &pool.Title, &pool.Description, &pool.CreatedAt, &pool.ExpiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Pool{}, errors.New("pool not found")
		}
		return models.Pool{}, err
	}

	// Query the options
	rows, err := p.db.Query(
		"SELECT id, text, votes FROM options WHERE pool_id = $1",
		id,
	)
	if err != nil {
		return models.Pool{}, err
	}
	defer rows.Close()

	// Populate the options
	pool.Options = []models.Option{}
	for rows.Next() {
		option := models.Option{}
		if err := rows.Scan(&option.ID, &option.Text, &option.Votes); err != nil {
			return models.Pool{}, err
		}
		pool.Options = append(pool.Options, option)
	}

	if err := rows.Err(); err != nil {
		return models.Pool{}, err
	}

	return pool, nil
}

func (p *PostgresDB) ListPools() []models.Pool {
	// Query all pools
	rows, err := p.db.Query(
		"SELECT id, title, description, created_at, expires_at FROM pools ORDER BY created_at DESC",
	)
	if err != nil {
		return []models.Pool{}
	}
	defer rows.Close()

	// Collect the pools
	pools := []models.Pool{}
	for rows.Next() {
		pool := models.Pool{}
		if err := rows.Scan(&pool.ID, &pool.Title, &pool.Description, &pool.CreatedAt, &pool.ExpiresAt); err != nil {
			return []models.Pool{}
		}

		// Get options for each pool
		optRows, err := p.db.Query(
			"SELECT id, text, votes FROM options WHERE pool_id = $1",
			pool.ID,
		)
		if err != nil {
			return []models.Pool{}
		}
		
		// Populate the options
		pool.Options = []models.Option{}
		for optRows.Next() {
			option := models.Option{}
			if err := optRows.Scan(&option.ID, &option.Text, &option.Votes); err != nil {
				optRows.Close()
				return []models.Pool{}
			}
			pool.Options = append(pool.Options, option)
		}
		optRows.Close()

		pools = append(pools, pool)
	}

	return pools
}

func (p *PostgresDB) UpdatePool(pool models.Pool) error {
	// Start a transaction
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update the pool
	_, err = tx.Exec(
		"UPDATE pools SET title = $1, description = $2, expires_at = $3 WHERE id = $4",
		pool.Title, pool.Description, pool.ExpiresAt, pool.ID,
	)
	if err != nil {
		return err
	}

	// Update the options
	for _, option := range pool.Options {
		_, err = tx.Exec(
			"UPDATE options SET text = $1, votes = $2 WHERE id = $3 AND pool_id = $4",
			option.Text, option.Votes, option.ID, pool.ID,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (p *PostgresDB) DeletePool(id string) error {
	// Cascading delete will handle the options
	_, err := p.db.Exec("DELETE FROM pools WHERE id = $1", id)
	return err
}

func (p *PostgresDB) CreateVote(vote models.Vote) error {
	// Start a transaction
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if user has already voted
	var count int
	err = tx.QueryRow(
		"SELECT COUNT(*) FROM votes WHERE pool_id = $1 AND user_id = $2",
		vote.PoolID, vote.UserID,
	).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("user already voted in this pool")
	}

	// Check if option exists
	var poolID string
	err = tx.QueryRow(
		"SELECT pool_id FROM options WHERE id = $1 AND pool_id = $2",
		vote.OptionID, vote.PoolID,
	).Scan(&poolID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("option not found in pool")
		}
		return err
	}

	// Insert the vote
	_, err = tx.Exec(
		"INSERT INTO votes (id, pool_id, option_id, user_id, created_at) VALUES ($1, $2, $3, $4, $5)",
		vote.ID, vote.PoolID, vote.OptionID, vote.UserID, vote.CreatedAt,
	)
	if err != nil {
		return err
	}

	// Increment the vote count
	_, err = tx.Exec(
		"UPDATE options SET votes = votes + 1 WHERE id = $1 AND pool_id = $2",
		vote.OptionID, vote.PoolID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (p *PostgresDB) GetVotesForPool(poolID string) ([]models.Vote, error) {
	// Query votes
	rows, err := p.db.Query(
		"SELECT id, pool_id, option_id, user_id, created_at FROM votes WHERE pool_id = $1",
		poolID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect the votes
	votes := []models.Vote{}
	for rows.Next() {
		vote := models.Vote{}
		if err := rows.Scan(&vote.ID, &vote.PoolID, &vote.OptionID, &vote.UserID, &vote.CreatedAt); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return votes, nil
}

func (p *PostgresDB) CreateUser(user models.User) error {
	// Insert the user
	_, err := p.db.Exec(
		"INSERT INTO users (id, username) VALUES ($1, $2)",
		user.ID, user.Username,
	)
	if err != nil {
		// Check for duplicate username
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				return errors.New("username already exists")
			}
		}
		return err
	}
	return nil
}

func (p *PostgresDB) GetUser(id string) (models.User, error) {
	// Query the user
	user := models.User{}
	err := p.db.QueryRow(
		"SELECT id, username FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}