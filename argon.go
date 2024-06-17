package argon

import (
	"crypto/rand"
	"log/slog"
	"math/big"

	"golang.org/x/crypto/argon2"
)

var (
	// SaltLen is the length of the salt.
	SaltLen = 64
	// Time is the number of iterations to use.
	Time uint32 = 1
	// Memory is the amount of memory used.
	Memory uint32 = 64 * 1024
	// Threads is the number of threads to use.
	Threads uint8 = 4
	// KeyLen is the length of the key to generate.
	KeyLen uint32 = 64
)

// Salt returns a random salt of length SaltLen.
//
//	salt, err := argon.Salt()
func Salt() ([]byte, error) {
	salt := make([]byte, 0, SaltLen)

	var err error
	var b *big.Int

	for range SaltLen {
		b, err = rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			slog.Error("error during the creation of a uniform random value", "err", err)
			return nil, err
		}
		salt = append(salt, byte(b.Int64()))
	}
	return salt, nil
}

// Argon2id returns the Argon2id hash of the password using the provided salt.
//
//	hash := argon.Argon2id("input", salt)
//
// without salt
//
//	hash := argon.Argon2id("input", nil)
func Argon2id(input string, salt []byte) []byte {
	return argon2.IDKey([]byte(input), salt, Time, Memory, Threads, KeyLen)
}

// Argon2idSalt returns the Argon2id hash of the password and a random salt.
//
//	hash, salt, err := argon.Argon2idSalt("password")
func Argon2idSalt(input string) ([]byte, []byte, error) {
	salt, err := Salt()
	if err != nil {
		return nil, nil, err
	}
	return argon2.IDKey([]byte(input), salt, Time, Memory, Threads, KeyLen), salt, nil
}
