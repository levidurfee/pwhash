package main

import (
    "fmt"
)

// formatKey formats the key
// '$argon2i$v=19$m=512,t=2,p=2$aI2R0hpDyLm3ltLa+1/rvQ$LqPKjd6n8yniKtAithoR7A'
func formatKey(key string, salt string) string {
    formatted := fmt.Sprintf("$argon2i$v=19$m=65536,t=1,p=4$%s$%s", salt, key)

    return formatted
}
