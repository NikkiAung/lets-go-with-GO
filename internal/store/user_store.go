package store

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type password struct {
	plainText *string
	hash []byte
}

func (p *password) Hash(plainText string) error {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(plainText), 12)
	if err != nil {
		return err
	}
	p.plainText = &plainText
	p.hash = hashpassword
	
	return nil
}

func (p *password) Match(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) : 		return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PasswordHash password `json:"password"`
	CreatedAt string `json:"created_at"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore {
		db : db,
	}
}

type UserStore interface {
	CreateUser(*User) error
	GetUserByUserName(username string) (*User, error)
	UpdateUser(*User) error
}

func (s *PostgresUserStore)CreateUser(user *User) error {
	query := `INSERT INTO users (username, email, passowrd_hash) VALUES ($1, $2, $3) RETURNING id, created_at`

	err := s.db.QueryRow(query,user.Username, user.Email, user.PasswordHash).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresUserStore)GetUserByUserName(username string) (*User, error) {
	user := &User {
		PasswordHash: password{},
	}

	query := `SELECT id, username, email, created_at FROM users WHERE username=$1`

	err := s.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresUserStore) UpdateUser(user *User) error {
	query := `UPDATE users SET username=$1, email=$2 WHERE id=$3`

	result, err := s.db.Exec(query, user.Username, user.Email, user.ID)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}