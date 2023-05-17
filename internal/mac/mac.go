// Package mac provides functions to interact with MAC addresses.
package mac

import (
	"bytes"
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

// Represents a vendor entity and its associated prefixes (as the IEEE database)
type vendor struct {
	name string
	prefixes [][]byte
}

// Print a MAC address in readable format
func (m *mac) toString() string {

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
func Validate(mac []byte) bool {
	return ValidAddress.Match(mac)
}

// Normalize convert a given MAC address into and uppercase format with a collon as separator.
func Normalize(mac []byte) ([]byte, error) {
	if !Validate(mac) {
		return []byte{}, fmt.Errorf("cannot normalize an invalid mac address")
	}

	// Replace "-" separator with ":"
	mac = bytes.ReplaceAll(mac, []byte{0x2d}, []byte{0x3a})

	// Replace lowercase characters to uppercase (yander but avoid needless conversions)
	mac = bytes.ReplaceAll(mac, []byte{0x61}, []byte{0x41})
	mac = bytes.ReplaceAll(mac, []byte{0x62}, []byte{0x42})
	mac = bytes.ReplaceAll(mac, []byte{0x63}, []byte{0x43})
	mac = bytes.ReplaceAll(mac, []byte{0x64}, []byte{0x44})
	mac = bytes.ReplaceAll(mac, []byte{0x65}, []byte{0x45})
	mac = bytes.ReplaceAll(mac, []byte{0x66}, []byte{0x46})
	mac = bytes.ReplaceAll(mac, []byte{0x67}, []byte{0x47})
	mac = bytes.ReplaceAll(mac, []byte{0x68}, []byte{0x48})
	mac = bytes.ReplaceAll(mac, []byte{0x69}, []byte{0x49})
	mac = bytes.ReplaceAll(mac, []byte{0x6a}, []byte{0x4a})
	mac = bytes.ReplaceAll(mac, []byte{0x6b}, []byte{0x4b})
	mac = bytes.ReplaceAll(mac, []byte{0x6c}, []byte{0x4c})
	mac = bytes.ReplaceAll(mac, []byte{0x6d}, []byte{0x4d})
	mac = bytes.ReplaceAll(mac, []byte{0x6e}, []byte{0x4e})
	mac = bytes.ReplaceAll(mac, []byte{0x6f}, []byte{0x4f})
	mac = bytes.ReplaceAll(mac, []byte{0x70}, []byte{0x50})
	mac = bytes.ReplaceAll(mac, []byte{0x71}, []byte{0x51})
	mac = bytes.ReplaceAll(mac, []byte{0x72}, []byte{0x52})
	mac = bytes.ReplaceAll(mac, []byte{0x73}, []byte{0x53})
	mac = bytes.ReplaceAll(mac, []byte{0x74}, []byte{0x54})
	mac = bytes.ReplaceAll(mac, []byte{0x75}, []byte{0x55})
	mac = bytes.ReplaceAll(mac, []byte{0x76}, []byte{0x56})
	mac = bytes.ReplaceAll(mac, []byte{0x77}, []byte{0x57})
	mac = bytes.ReplaceAll(mac, []byte{0x78}, []byte{0x58})
	mac = bytes.ReplaceAll(mac, []byte{0x79}, []byte{0x59})
	mac = bytes.ReplaceAll(mac, []byte{0x7a}, []byte{0x5a})

	return mac, nil
}

func Rand() (string, error) {
	randBytes := make([]byte, 6)

	_, err := rand.Read(randBytes)

	if err != nil {
		return "", err
	}

	randMac := fmt.Sprintf("%x", randBytes[0])

	for i := 1; i < 6; i++ {
		randMac = strings.Join([]string{randMac, fmt.Sprintf("%02x", randBytes[i])}, ":")
	}

	return randMac, nil
}

func FromVendor(name string) {
	apple := vendor{
		name: "Apple, Inc.",
		prefixes: [][]byte{{0x00, 0x03, 0x93}, {0x00, 0x0a, 0x27}},
	}
}