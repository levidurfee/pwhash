package main

import (
    "golang.org/x/crypto/argon2"
)

func makeArgon(password string, salt string) string {
    decodedSalt, _ := b64d(salt)
    hash := argon2.Key([]byte(password), decodedSalt, 1, 64*1024, 4, 33)

    return b64e(hash)
}
