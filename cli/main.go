/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/gotbots/autobots/internal/cmd"
	"github.com/gotbots/autobots/internal/cmd/shared"
)

func main() {
	deps := shared.Dependencies{}
	rootCmd := cmd.NewRoot(deps)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
