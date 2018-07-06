package main

import (
    "flag"
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    passwordPtr := flag.String("password", "", "Password to hash")
    costPtr := flag.Int("cost", 10, "BCrypt Cost")
    
    flag.Parse()
    
    // fmt.Println(*passwordPtr)
    // fmt.Println(*costPtr)

    password := []byte(*passwordPtr)

    hashedPassword, err := bcrypt.GenerateFromPassword(password, *costPtr)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(hashedPassword))
}