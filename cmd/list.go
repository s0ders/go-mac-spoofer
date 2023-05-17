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
	Short:   "List availables NIC and their MAC address",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		nics, err := nic.List()

		if err != nil {
			fmt.Printf("failed to list interfaces: %s\n", err)
			os.Exit(-1)
		}

		fmt.Printf("%-17s %-17s\n\n", "NAME", "ADDRESS")
		for _, nic := range nics {
			fmt.Printf("%-17s %-17s\n", nic.Name, nic.Address)
		}
	},
}
