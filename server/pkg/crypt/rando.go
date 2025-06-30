package crypt

import (
	"crypto/rand"
	"math/big"
)

var CharList = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]{}<>?,./")

// GenerateRandomStringFromLetters generates a random string of the given length using the provided runes
// this function panics if
// - the provided length is less than 1
// - if the provided runes are empty
// - if os fails to read random bytes
func GenerateRandomStringFromLetters(length int, letterRunes []rune) string {
	b := make([]rune, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			panic(err)
		}
		b[i] = letterRunes[num.Int64()]
	}
	return string(b)
}

func GererateRandomInt(min, max int64) (int64, error) {
    n := max - min + 1
	
    x, err := rand.Int(rand.Reader, big.NewInt(n))

    if err != nil {
        return 0, err
    }

    return x.Int64() + min, nil
}

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
