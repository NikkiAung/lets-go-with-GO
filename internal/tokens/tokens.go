package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

type Token struct {
    Plaintext string    `json:"token"`
    Hash      []byte    `json:"-"`
    UserId    int       `json:"-"`
    Expiry    time.Time `json:"expiry"`
    Scope     string    `json:"-"`
}

const (
    ScopeAuth = "authentication"
)

func GenerateToken(userID int, ttl time.Duration, scope string) (*Token, error) {
    token := &Token{
        UserId:  userID,
        Expiry: time.Now().Add(ttl),
        Scope:  scope,
    }

	// bytes -> base32 string -> sha

    emptyBytes := make([]byte, 32)
    _, err := rand.Read(emptyBytes)
    if err != nil {
        return nil, err
    }

	// you are generating a random token string,
	// but the randomness comes from crypto/rand,
	// while Base32 encodes it, and SHA-256 secures it.

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(emptyBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}
