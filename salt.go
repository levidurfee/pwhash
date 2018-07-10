package main

import (
    "crypto/rand"
)

func salt(len uint8) (string, error) {
    salt := make([]byte, len)
    _, err := rand.Read(salt)

    return b64e(salt), err
}
