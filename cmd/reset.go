package cmd

import (
	"fmt"

	"github.com/s0ders/go-mac-spoofer/internal/nic"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset INTERFACE",
	Short: "Attempts to change a network interface MAC address back to its original value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := nic.ResetMAC(name)

		if err != nil {
			fmt.Printf("failed to reset MAC: %s", err)
		}

		fmt.Printf("successfully reset %s MAC address\n", string(name))
	},
}
