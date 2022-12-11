package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
	switch os.Args[1] {
	case "hash":
		//hash the password
		hash(os.Args[2])

	case "compare":
		//compare the password and the hash
		compare(os.Args[2], os.Args[3])
	default:
		//invalid command
		fmt.Printf("Invalid command: %q\n", os.Args[1])
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", password)
	}
	fmt.Printf("Hashed value the password %q is %v\n", password, string(hashedBytes))
}

func compare(password, hash string) {
	fmt.Printf("TODO: Compare the password %q and the hash %q\n", password, hash)
}
