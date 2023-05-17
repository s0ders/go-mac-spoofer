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
	spoofCmd.Flags().BoolVarP(&generateRandomAddress, "random", "r", false, "Use a pseudo-randomly generated MAC addresses")
	spoofCmd.Flags().StringVarP(&newMACAddress, "address", "a", "01:23:45:67:89:10", "The MAC address to spoof")

	rootCmd.AddCommand(spoofCmd)
}

var spoofCmd = &cobra.Command{
	Use:     "spoof INTERFACE",
	Aliases: []string{"change"},
	Short:   "Attempts to spoof a network interface MAC address",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		nicName := args[0]
		var newAddr string

		if generateRandomAddress || len(newMACAddress) == 0 {
			randMAC, err := mac.Rand()

			newAddr = randMAC

			if err != nil {
				log.Fatalf("could not generate random address: %s\n", err)
			}
		} else {
			newAddr = newMACAddress
		}

		err := nic.ChangeMAC(nicName, newAddr)

		if err != nil {
			log.Fatalf("failed to change address: %s\n", err)
		}

		fmt.Printf("successfully changed %s address to %s\n", string(nicName), string(newAddr))

	},
}
