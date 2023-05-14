package main

import (
	"log"
	"os"
	"fmt"

	"github.com/s0ders/go-mac-spoofer/internal/nic"

	"github.com/spf13/cobra"
)

func main() {

	// TODO: implement the following commands:
	// - "list", list all NIC and their MAC addresses
	// - "change", change a NIC mac address, either from a given address, a random "-random" or random with specific vendor "-vendor"
	// - "reset", reset a NIC address to its factory address

	var rootCmd = &cobra.Command{
		Use:   "mac-spoofer",
		Short: "mac-spoofer - a simple CLI to spoof MAC address",
		Long:  "mac-spoofer - open source CLI built in Go to spoof a NIC MAC address on Windows, MacOS and Linux",
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

	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("There was an error while executing CLI: %s", err)
		os.Exit(-1)
	}

}
