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
	Short: "Attempts to change a NIC MAC address back to its factory value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := []byte(args[0])

		err := nic.ResetMAC(name)

		if err != nil {
			fmt.Printf("failed to reset MAC: %s", err)
		}

		fmt.Printf("Successfully reset %s MAC address\n", string(name))
	},
}
