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
		Long:  "go-mac-spoofer - open source CLI built in Go to spoof a NIC MAC address on Windows, MacOS and Linux",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
