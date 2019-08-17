package main

import (
	"os"
)

var version = "0.0.1"

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
