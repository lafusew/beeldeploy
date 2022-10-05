package main

import (
	"log"
	"os"
	"os/exec"
)

// Check the installation of gcloud CLI application
func CheckInstallCLI() error {
	cmd := exec.Command("gcloud", "--version")
	err := cmd.Run()

	return err
}

// Executes the gcloud cmd and print the output
func Exec(args []string) {
	cmd := exec.Command("gcloud", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Cloud function deploy failed")
	}
}
