package cmd

import (
	"fmt"
	"log"

	"github.com/s0ders/go-mac-spoofer/internal/mac"
	"github.com/s0ders/go-mac-spoofer/internal/nic"
	"github.com/spf13/cobra"
)

var (
	generateRandomAddress bool
	newMACAddress         string
)

func init() {
	changeCmd.Flags().BoolVarP(&generateRandomAddress, "random", "r", false, "Use a pseudo-randomly generated MAC addresses")
	changeCmd.Flags().StringVarP(&newMACAddress, "address", "a", "01:23:45:67:89:10", "The MAC address to spoof")

	rootCmd.AddCommand(changeCmd)
}

var changeCmd = &cobra.Command{
	Use:     "spoof INTERFACE",
	Aliases: []string{"change"},
	Short:   "Attempts to spoof an NIC MAC address",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		nicName := args[0]
		var newAddr string

		if generateRandomAddress || len(newMACAddress) == 0 {
			randMAC, err := mac.Rand()

			newAddr = randMAC

			if err != nil {
				log.Fatalf("could not generate random address: %s", err)
			}
		} else {
			newAddr = newMACAddress
		}

		err := nic.ChangeMAC(nicName, newAddr)

		if err != nil {
			log.Fatalf("failed to change address: %s\n", err)
		}

		fmt.Printf("successfully changed %s address to %s", string(nicName), string(newAddr))

	},
}
