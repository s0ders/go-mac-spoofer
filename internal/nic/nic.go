// Package nic provides functions to interact with the system NICs
package nic

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	"github.com/s0ders/go-mac-spoofer/internal/mac"
)

var (
	PATH_TO_AIRPORT = "/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport"
	PLATFORM        = runtime.GOOS
)

type NIC struct {
	Name    []byte
	Address []byte
}

// List returns a slice with the name and address of each available NICs.
func List() ([]NIC, error) {
	nics := make([]NIC, 0)

	switch PLATFORM {
	case "darwin":
		NICNameRegex := regexp.MustCompile(`Device: ([a-z0-9]+)`)
		NICAddressRegex := regexp.MustCompile(`Ethernet Address: ([a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2})`)

		cmd, err := exec.Command("networksetup", "-listallhardwareports").Output()

		if err != nil {
			return []NIC{}, err
		}

		nameMatches := NICNameRegex.FindAllSubmatch(cmd, -1)
		addressMatches := NICAddressRegex.FindAllSubmatch(cmd, -1)

		for i := range nameMatches {
			name := nameMatches[i][1]
			address, err := mac.Normalize(addressMatches[i][1])

			if err != nil {
				return nil, fmt.Errorf("could not normalize NIC address: %w", err)
			}

			nics = append(nics, NIC{name, address})
		}
	case "linux":

	default:
		return []NIC{}, fmt.Errorf("unknown platform: %s", PLATFORM)
	}

	return nics, nil
}

// Exists checks if a given NIC exists on the system.
func Exists(nicName []byte) (bool, error) {
	existing, err := List()

	if err != nil {
		return false, err
	}

	for _, existingNic := range existing {
		if string(existingNic.Name) == string(nicName) {
			return true, nil
		}
	}

	return false, nil
}

// ChangeMAC attempts to change a NIC MAC address.
func ChangeMAC(nicName, newMAC []byte) error {
	var err error

	if !mac.Validate(newMAC) {
		return fmt.Errorf("new MAC address \"%s\" is invalid", newMAC)
	}

	newMAC, err = mac.Normalize(newMAC)

	if err != nil {
		return err
	}

	switch PLATFORM {
	case "darwin":

		if os.Geteuid() != 0 {
			log.Fatalf("This program must be executed as root (UID 0) to be able to change network card interface settings")
		}

		// Disassociate from wi-fi network without turning off wi-fi or the device
		_, err := exec.Command(PATH_TO_AIRPORT, "-z").Output()

		if err != nil {
			log.Fatalf("error happened while trying to invoke airport: %s", err.Error())
		}

		// Changing MAC address
		_, err = exec.Command("ifconfig", string(nicName), "ether", string(newMAC)).Output()

		if err != nil {
			log.Fatalf("error happened while trying to invoke ifconfig: %s", err.Error())
		}

		// Restart airport on device to reassociate with known networks
		_, err = exec.Command("networksetup", "-setairportpower", string(nicName), "off").Output()

		if err != nil {
			log.Fatalf("error happened while trying to set airport power off: %s", err.Error())
		}

		_, err = exec.Command("networksetup", "-setairportpower", string(nicName), "on").Output()

		if err != nil {
			log.Fatalf("error happened while trying to set airport power on: %s", err)
		}

		return nil
	}

	return nil
}

// ResetMAC attempts to reset a NIC MAC address to its factory value
func ResetMAC(name []byte) error {

	exists, err := Exists(name)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("could not find a NIC named %s", string(name))
	}

	switch PLATFORM {
	case "darwin":
		nics, err := List()

		if err != nil {
			return err
		}

		var originalMAC []byte

		for _, nic := range nics {
			if string(nic.Name) == string(name) {
				originalMAC = nic.Address
			}
		}

		err = ChangeMAC(name, originalMAC)

		if err != nil {
			return err
		}
	}

	return nil
}
