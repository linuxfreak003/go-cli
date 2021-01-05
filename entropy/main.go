package main

import (
	"fmt"
	"os"
	"strings"

	validator "github.com/lane-c-wagner/go-password-validator"
)

func main() {
	pass := strings.Join(os.Args[1:], " ")
	entropy := validator.GetEntropy(pass)
	fmt.Printf("Entropy: %0.2f bits\n", entropy)
}
