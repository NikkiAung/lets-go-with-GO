package store

import (
	"database/sql"
	"time"

	"github.com/NikkiAung/go-fundmentals/internal/tokens"
)

type PostgresTokenStore struct {
    db *sql.DB
}

func NewPostgresTokenStore(db *sql.DB) *PostgresTokenStore {
    return &PostgresTokenStore{
        db: db,
    }
}

type TokenStore interface {
    CreateToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error)
	SaveToken(token *tokens.Token) error
	DeleteTokensByUserId(userID int, scope string) error
}

func (s *PostgresTokenStore) CreateToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error) {
	token, err := tokens.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = s.SaveToken(token)

	return token, err
}

func (s *PostgresTokenStore) SaveToken(token *tokens.Token) error {
	query := `INSERT INTO tokens (hash,user_id,expiry,scope) VALUES ($1,$2,$3,$4)`
	_, err := s.db.Exec(query,token.Hash, token.UserId, token.Expiry, token.Scope)

	return err
}

func (s *PostgresTokenStore) DeleteTokensByUserId(userID int, scope string) error {
    query := `DELETE FROM tokens WHERE scope = $1 AND user_id = $2`

    _, err := s.db.Exec(query, scope, userID)
    return err
}


