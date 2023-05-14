// Package nic provides functions to interact with the system NICs
package nic

import (
	"log"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
)

var (
	PATH_TO_AIRPORT = "/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport"
)

type NIC struct {
	Name    []byte
	Address []byte
}

// List() returns a slice with the name of each available NICs.
func List() ([]NIC, error) {
	platform := runtime.GOOS

	nics := make([]NIC, 0)

	switch platform {
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
			nics = append(nics, NIC{nameMatches[i][1], addressMatches[i][1]})
		}
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
		if reflect.DeepEqual(existingNic, nicName) {
			return true, nil
		}
	}

	return false, nil
}

// ChangeMAC attemps to change a NIC MAC address.
func ChangeMAC(nicName, newMac []byte) error {
	platform := runtime.GOOS

	switch platform {
	case "darwin":

		if os.Geteuid() != 0 {
			log.Fatalf("Are you running this program as root ?")
		}

		// Disassociate from wi-fi network without turning off wi-fi or the device
		_, err := exec.Command(PATH_TO_AIRPORT, "-z").Output()

		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}

		// Changing MAC address
		_, err = exec.Command("ifconfig", string(nicName), "ether", string(newMac)).Output()

		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}

		// Restart airport on device to reassociate with known networks
		_, err = exec.Command("networksetup", "-setairportpower", string(nicName), "off").Output()

		if err != nil {
			log.Fatalf("error: %s", err.Error())
		}

		_, err = exec.Command("networksetup", "-setairportpower", string(nicName), "on").Output()

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
