package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := myFunction()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = myFunction2()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func myFunction() (string, error) {
	// …
	failureCondition := true
	if failureCondition {
		return "", errors.New("An error occurred")
	}
	// …
	return "Success", nil
}

// wrapping errors
func myFunction2() (string, error) {
	// …
	failureCondition := true
	someOtherError := errors.New("An error occurred")
	if failureCondition {
		return "", fmt.Errorf("An error occurred: %w", someOtherError)
	}
	// …
	return "Success", nil
}

// Custom Errors
type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}
