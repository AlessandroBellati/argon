package argon_test

import (
	"log"
	"testing"

	"github.com/AlessandroBellati/argon"
)

func Benchmark_Salt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, err := argon.Salt()
		if err != nil {
			b.Error(err)
		}
		log.Println(s)
	}
}

func Test_Argon2id(t *testing.T) {
	salt, err := argon.Salt()
	if err != nil {
		t.Error(err)
	}
	hash := argon.Argon2id("password", salt)
	log.Println(hash)
}
