package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "mga",
	RunE: rootCmdRun,
}

func rootCmdRun(cmd *cobra.Command, _ []string) error {
	if err := cmd.Help(); err != nil {
		return fmt.Errorf("error printing help: %v", err)
	}
	return nil
}
