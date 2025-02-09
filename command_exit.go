package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Printf("Ending session...\n")
	os.Exit(0)
	return nil
}
