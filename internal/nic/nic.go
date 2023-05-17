// Package nic provides functions to interact with the system NICs
package nic

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/s0ders/go-mac-spoofer/internal/mac"
)

var (
	PATH_TO_AIRPORT = "/System/Library/PrivateFrameworks/Apple80211.framework/Resources/airport"
	PLATFORM        = runtime.GOOS
)

type NetworkInterface struct {
	Name    string
	Address string
}

// List returns a slice with the name and address of each available NICs.
func List() ([]NetworkInterface, error) {
	nics := make([]NetworkInterface, 0)

	switch PLATFORM {
	case "darwin":
		NICNameRegex := regexp.MustCompile(`Device: ([a-z0-9]+)`)
		NICAddressRegex := regexp.MustCompile(`Ethernet Address: ([a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2}\:[a-z0-9]{2})`)

		// List available interfaces and their MAC addresses
		cmd, err := exec.Command("networksetup", "-listallhardwareports").Output()

		if err != nil {
			return nil, err
		}

		nameMatches := NICNameRegex.FindAllStringSubmatch(string(cmd), -1)
		addressMatches := NICAddressRegex.FindAllStringSubmatch(string(cmd), -1)

		for i := range nameMatches {
			name := nameMatches[i][1]
			address, err := mac.Normalize(addressMatches[i][1])

			if err != nil {
				return nil, fmt.Errorf("could not normalize NIC address: %w", err)
			}

			nics = append(nics, NetworkInterface{name, address})
		}
	case "linux":
		nicNames := make([]string, 0)

		// Get all interfaces available
		filepath.Walk("/sys/class/net/", func(path string, info fs.FileInfo, err error) error {
			nicNames = append(nicNames, info.Name())
			return nil
		})

		// Get MAC address for each available interface
		for _, name := range nicNames {

			if name == "net" {
				continue
			}

			addrPath := "/sys/class/net/" + name + "/address"

			if _, err := os.Stat(addrPath); errors.Is(err, os.ErrNotExist) {
				return nil, fmt.Errorf("could not get address for %s", name)
			}

			addr, err := os.ReadFile(addrPath)

			if err != nil {
				return nil, fmt.Errorf("could not get address for %s", name)
			}

			normalizedAddr, err := mac.Normalize(string(addr))

			if err != nil {
				return nil, fmt.Errorf("could not normalize address for interface %s", name)
			}

			nics = append(nics, NetworkInterface{name, normalizedAddr})
		}
	default:
		return []NetworkInterface{}, fmt.Errorf("unknown platform: %s", PLATFORM)
	}

	return nics, nil
}

// Exists checks if a given NIC exists on the system.
func Exists(nicName string) (bool, error) {
	existing, err := List()

	if err != nil {
		return false, err
	}

	for _, existingNic := range existing {
		if existingNic.Name == nicName {
			return true, nil
		}
	}

	return false, nil
}

// ChangeMAC attempts to change a network interface MAC address.
func ChangeMAC(interfaceName, newAddr string) error {
	var err error

	if !mac.Validate(newAddr) {
		return fmt.Errorf("new address \"%s\" is invalid", newAddr)
	}

	newAddr, err = mac.Normalize(newAddr)

	if err != nil {
		return err
	}

	switch PLATFORM {
	case "darwin":

		if os.Geteuid() != 0 {
			log.Fatalf("This program must be executed as root (UID 0) to be able to change network card interface settings")
		}

		if _, err := os.Stat(PATH_TO_AIRPORT); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("cannot change address on macOS without airport")
		}

		// Disassociate from wi-fi network without turning off wi-fi or the device
		_, err := exec.Command(PATH_TO_AIRPORT, "-z").Output()

		if err != nil {
			log.Fatalf("error happened while trying to invoke airport: %s", err.Error())
		}

		// Changing MAC address
		_, err = exec.Command("ifconfig", interfaceName, "ether", newAddr).Output()

		if err != nil {
			log.Fatalf("error happened while trying to invoke ifconfig: %s", err.Error())
		}

		// Restart airport on device to reassociate with known networks
		_, err = exec.Command("networksetup", "-setairportpower", interfaceName, "off").Output()

		if err != nil {
			log.Fatalf("error happened while trying to set airport power off: %s", err.Error())
		}

		_, err = exec.Command("networksetup", "-setairportpower", interfaceName, "on").Output()

		if err != nil {
			log.Fatalf("error happened while trying to set airport power on: %s", err)
		}

		return nil
	}

	return nil
}

// ResetMAC attempts to reset a NIC MAC address to its factory value
func ResetMAC(name string) error {

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

		var originalAddr string

		for _, nic := range nics {
			if nic.Name == name {
				originalAddr = nic.Address
			}
		}

		err = ChangeMAC(name, originalAddr)

		if err != nil {
			return err
		}
	}

	return nil
}
