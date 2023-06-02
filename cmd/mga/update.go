package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streamingfast/dstore"
	"github.com/streamingfast/hivemapper-mga-tool/mga"
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
	data, h, err := mga.Download(args[0])
	if err != nil {
		return fmt.Errorf("error downloading MGA data: %w", err)
	}

	store, err := dstore.NewStore(args[1], "", "", true)
	if err != nil {
		return fmt.Errorf("error creating store %q: %w", args[1], err)
	}

	err = writeToS3(cmd.Context(), store, data, "mgaoffline.ubx")
	if err != nil {
		return fmt.Errorf("writing MGA data file %s: %w", "mgaoffline.ubx", err)
	}

	err = writeToS3(cmd.Context(), store, []byte(h), "mgaoffline.md5")
	if err != nil {
		return fmt.Errorf("writing MGA data file %s: %w", "mgaoffline.md5", err)
	}

	//remove the later ...
	err = writeToS3(cmd.Context(), store, data, "mgaoffine.ubx")
	if err != nil {
		return fmt.Errorf("writing MGA data file %s: %w", "mgaoffine.ubx", err)
	}

	err = writeToS3(cmd.Context(), store, []byte(h), "mgaoffine.md5")
	if err != nil {
		return fmt.Errorf("writing MGA data file %s: %w", "mgaoffine.md5", err)
	}

	fmt.Println("All done!")
	return nil
}

func writeToS3(ctx context.Context, store dstore.Store, data []byte, filename string) error {
	fmt.Println("Writing", filename, "with size:", len(data))
	buf := &bytes.Buffer{}
	_, err := buf.Write(data)
	if err != nil {
		return fmt.Errorf("creating buffer: %w", err)
	}

	err = store.WriteObject(ctx, filename, buf)
	if err != nil {
		return fmt.Errorf("writing MGA data file %s: %w", filename, err)
	}

	return nil
}
