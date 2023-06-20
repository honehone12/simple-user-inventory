package utils

import (
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

const (
	N            = 1 << 0xf
	P            = 16
	Q            = 4
	HashLength32 = 32
)

type PasswordHasher struct {
	n          int
	p          int
	q          int
	hashLength int
	password   []byte
}

type HashedPassword struct {
	Salt []byte
	DK   []byte
}

func NewPasswordHasher(password string) *PasswordHasher {
	return &PasswordHasher{
		n:          N,
		p:          P,
		q:          Q,
		hashLength: HashLength32,
		password:   []byte(password),
	}
}

func (ph *PasswordHasher) Hash() (*HashedPassword, error) {
	salt := make([]byte, HashLength32)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	dk, err := scrypt.Key(
		ph.password,
		salt,
		ph.n, ph.p, ph.q,
		ph.hashLength,
	)
	if err != nil {
		return nil, err
	}

	return &HashedPassword{
		Salt: salt,
		DK:   dk,
	}, nil
}
