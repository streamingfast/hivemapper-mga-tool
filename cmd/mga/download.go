package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streamingfast/hivemapper-mga-tool/mga"
	"os"
)

var downloadCmd = &cobra.Command{
	Use:   "download [mga-data-url]",
	RunE:  downloadCmdRun,
	Short: "Download MGA offline data from u-blox and compute MD5 checksum",
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func downloadCmdRun(_ *cobra.Command, args []string) error {
	fmt.Println("Downloading MGA offline data")
	data, h, err := mga.Download(args[0])
	if err != nil {
		return fmt.Errorf("error downloading MGA data: %w", err)
	}

	err = os.WriteFile("mgaoffline.ubx", data, 0644)
	if err != nil {
		return fmt.Errorf("error writing MGA data file %s: %w", "mgaoffline.ubx", err)
	}

	err = os.WriteFile("mgaoffline.md5", []byte(h), 0644)
	if err != nil {
		return fmt.Errorf("error writing MGA md5 file %s: %w", "mgaoffline.md5", err)
	}
	fmt.Println("MGA data downloaded", h)
	return nil
}
