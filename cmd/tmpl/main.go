package main

import (
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
