package main

import (
	"fmt"
	"github.com/Mrzrb/astra/cli/astra/cmd"
	"os"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err.Error())
		os.Exit(1)
	}
}
