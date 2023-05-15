package cmd

import (
	"fmt"
	"os"

	"github.com/s0ders/go-mac-spoofer/internal/nic"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List each available NIC and its MAC address",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		nics, err := nic.List()

		if err != nil {
			fmt.Printf("Could not list NICs")
			os.Exit(-1)
		}

		fmt.Printf("%-10s %-17s\n\n", "Name", "Address")
		for _, nic := range nics {
			fmt.Printf("%-10s %-17s\n", nic.Name, nic.Address)
		}
	},
}