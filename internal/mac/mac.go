// Package mac provides functions to interact with MAC addresses.
package mac

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
)

var (
	ValidAddress = regexp.MustCompile(`^([0-9A-Fa-f]{2}[\:\-]{1}){5}[0-9A-Fa-f]{2}$`)
)

// Represents a MAC address, alias for [6]byte
type mac [6]byte

// Print a MAC address in readable format
func (m mac) toString() string {

	macToString := byteToString(m[0])

	for i := 1; i < 6; i++ {
		macToString += ":" + byteToString(m[i])
	}

	return macToString
}

// Print a byte hex value and strips "0x" prefix
func byteToString(b byte) string {
	hexString := fmt.Sprintf("%02x", b)

	return strings.Replace(hexString, "0x", "", 1)
}

// Validate checks if the given slice of bytes if a valid MAC address.
func Validate(mac string) bool {
	return ValidAddress.MatchString(mac)
}

// Normalize convert a given MAC address into and uppercase format with a collon as separator.
func Normalize(mac string) (string, error) {
	if !Validate(mac) {
		return "", fmt.Errorf("cannot normalize an invalid mac address")
	}

	// Replace "-" separator with ":"
	mac = strings.ReplaceAll(mac, "-", ":")

	// Replace lowercase characters to uppercase
	for i := 97; i < 122; i++ {
		mac = strings.ReplaceAll(mac, fmt.Sprintf("%c", rune(i)), fmt.Sprintf("%c", rune(i-32)))
	}

	return mac, nil
}

func Rand() (string, error) {
	randBytes := make([]byte, 6)

	_, err := rand.Read(randBytes)

	if err != nil {
		return "", err
	}

	randAddress := mac(randBytes)

	return randAddress.toString(), nil
}
