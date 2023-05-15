package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "mac-spoofer",
		Short: "mac-spoofer - a simple CLI to spoof MAC address",
		Long:  "mac-spoofer - open source CLI built in Go to spoof a NIC MAC address on Windows, MacOS and Linux",
	}
)

func Execute() error {
	return rootCmd.Execute()
}