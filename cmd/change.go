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

		nicName := []byte(args[0])
		var newMAC []byte

		if generateRandomAddress || len(newMACAddress) == 0 {
			randMAC, err := mac.Rand()

			newMAC = []byte(randMAC)

			if err != nil {
				log.Fatalf("could not generate random address: %s", err)
			}
		} else {
			newMAC = []byte(newMACAddress)
		}

		err := nic.ChangeMAC(nicName, newMAC)

		if err != nil {
			log.Fatalf("failed to change MAC: %s", err)
		}

		fmt.Printf("Successfully changed %s MAC to %s", string(nicName), string(newMAC))

	},
}
