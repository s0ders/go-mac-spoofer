package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version string
	rootCmd = &cobra.Command{
		Use:   "go-mac-spoofer",
		Version: version,
		Short: "go-mac-spoofer - a simple CLI to spoof MAC address",
		Long:  "go-mac-spoofer - open source CLI to spoof network interfaces MAC addresses on Windows, macOS and Linux",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
