package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [mga-data-url] [s3-bucket]",
	RunE:  updateCmdRun,
	Short: "Download MGA offline data from u-blox service and update it on S3",
	Args:  cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateCmdRun(cmd *cobra.Command, args []string) error {
	fmt.Println("Updating MGA offline data")

	return fmt.Errorf("not implemented")
}
