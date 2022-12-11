package main

import (
	"fmt"
	"os"
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
	fmt.Printf("TODO: Hash the password %q\n", password)
}

func compare(password, hash string) {
	fmt.Printf("TODO: Compare the password %q and the hash %q\n", password, hash)
}
