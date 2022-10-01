package main

import (
	"fmt"
)

// Check the installation of gcloud CLI application
func CheckInstallCLI() error {
	// fmt.Printf("Check gcloud installation and auth\n\n")
	return nil
}

// Executes the cmd and print the output
func Exec(cmd string) {
	fmt.Printf("Executes cmd:\n\n%s\n", cmd)
}
