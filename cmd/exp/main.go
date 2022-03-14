package main

import (
	"errors"
	"fmt"
)

const (
	badInput  = "abc"
	goodInput = "xyz"
)

//Is Error function setup...

// var ErrBadInput = errors.New("bad input")

// func validateInput(input string) error {
// 	if input == badInput {
// 		return fmt.Errorf("validateInput: %w", ErrBadInput)
// 	}
// 	return nil
// }

// func validateInput(input string) error {
// 	if input == badInput {
// 		return fmt.Errorf("validateInput: %w", ErrBadInput)
// 	}
// 	return nil
// }

// func main() {
// 	//input := badInput
// 	input := goodInput

// 	err := validateInput(input)
// 	if errors.Is(err, ErrBadInput) {
// 		fmt.Println("bad input error")
// 	} else {
// 		fmt.Println("no input error")
// 	}
// }

// As Error function setup
type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
	}
	return nil
}

func main() {
	// input := badInput
	input := goodInput

	err := validateInput(input)
	var badInputErr *BadInputError
	if errors.As(err, &badInputErr) {
		fmt.Println("bad input error occured:", badInputErr)
	} else {
		fmt.Println("no input error")
	}
}
