package main

import (
    "encoding/base64"
)

// b64e encodes []byte into base64
func b64e(in []byte) string {
    return base64.StdEncoding.EncodeToString(in)
}

// b64d decodes string into []byte
func b64d(in string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(in)
}
