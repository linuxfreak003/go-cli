package main

import (
	"fmt"

	validator "github.com/lane-c-wagner/go-password-validator"
)

func main() {
	entropy := validator.GetEntropy("pass")
	fmt.Printf("Entropy: %f\n", entropy)
}
