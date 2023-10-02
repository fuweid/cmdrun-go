package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fatal(run())
}

func run() error {
	if err := validateArgs(); err != nil {
		return err
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start %s: %w", os.Args[1], err)
	}
	return nil
}

func validateArgs() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("command is required")
	}

	cmdName := os.Args[1]
	_, err := exec.LookPath(cmdName)
	if err != nil {
		return fmt.Errorf("failed to ensure %s: %w", cmdName, err)
	}
	return nil
}

func fatal(err error) {
	if err == nil {
		return
	}

	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
