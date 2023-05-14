// Package nic provides functions to interact with the system NICs
package nic

import (
	"fmt"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
)

var (
	PATH_TO_AIRPORT = "/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport"
)

// ListNICs() returns a slice with the name of each available NICs.
func ListNICs() ([][]byte, error) {
	platform := runtime.GOOS

	nics := make([][]byte, 0)

	switch platform {
	case "darwin":
		darwinNICRegex := regexp.MustCompile(`Device: ([a-z0-9]+)`)

		cmd, err := exec.Command("networksetup", "-listallhardwareports").Output()

		if err != nil {
			return [][]byte{}, err
		}

		matches := darwinNICRegex.FindAllSubmatch(cmd, -1)
		for _, match := range matches {
			nics = append(nics, match[1])
		}
	}

	return nics, nil
}

// Exists checks if a given NIC exists on the system.
func Exists(nicName []byte) (bool, error) {
	existing, err := ListNICs()

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

		// Disassociate from wi-fi network without turning off wi-fi or the device
		_, err := exec.Command(PATH_TO_AIRPORT, "-z").Output()

		if err != nil {
			return err
		}

		// Changing MAC address
		ifconfigCmd := fmt.Sprintf("ifconfig %s ether %s", string(nicName), string(newMac))
		_, err = exec.Command(ifconfigCmd).Output()

		if err != nil {
			return err
		}

		// Restart airport on device to reassociate with known networks
		_, err = exec.Command("networksetup", "-setairportpower", string(nicName), "off").Output()

		if err != nil {
			return err
		}

		_, err = exec.Command("networksetup", "-setairportpower", string(nicName), "on").Output()

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}