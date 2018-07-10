package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Starting...")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.Trim(password, " \n")

	salt, _ := salt(18)
	hash := makeArgon(password, salt)
	key := formatKey(hash, salt)

	fmt.Println("=======================\n\nARGON")
	fmt.Println(key)
	fmt.Println("\n=======================\n")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 11)

	fmt.Println("BCRYPT")
	fmt.Println(string(hashedPassword))
}
