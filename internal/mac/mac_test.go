package mac

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type table struct {
		mac []byte
		isValid bool
	}

	tests := []table{
		{[]byte("00-00-00-00-00-00"), true},
		{[]byte("00:00:00:00:00:00"), true},
		{[]byte("aa:10:89:ad:00:ff"), true},
		{[]byte("FF:AB:CD:E4:80:90"), true},
		{[]byte("FF-AB-CD-E4-80-90"), true},
		{[]byte("fg:00:00:00:00:00"), false}, 		// out of range
		{[]byte("ZZ-AB-CD-E4-80-90"), false},		// out of range
		{[]byte("PM:AG:CD:E4:80:90"), false},		// out of range
		{[]byte("aa:aa:aa:aa:aa"), false},			// too short
		{[]byte("aa:aa:aa:aa:aa:aa:aa"), false},	// too long
	}

	for _, test := range tests {
		got := Validate(test.mac)
		if got != test.isValid {
			t.Errorf("got: %t, want: %t with @ %s", got, test.isValid, string(test.mac))
		}
	}
}

func TestNormalize(t *testing.T) {
	type table struct {
		mac []byte
		expect []byte
	}

	tests := []table{
		{[]byte("aa-bc-dc-12-12-12"), []byte("AA:BC:DC:12:12:12")},
		{[]byte("aa:bc:dc:12:12:12"), []byte("AA:BC:DC:12:12:12")},
		{[]byte("AA:BC:DC:12:12:12"), []byte("AA:BC:DC:12:12:12")},
	}

	for _, test := range tests {
		got, err := Normalize(test.mac)

		if err != nil {
			t.Errorf("error while normalizing: %s", err)
		}

		if string(got) != string(test.expect) {
			t.Errorf("got: %s, want: %s", string(got), string(test.expect))
		}
	}
}

func TestRand(t *testing.T) {
	randMAC, err := Rand()

	if err != nil {
		t.Errorf("error while generating random MAC: %s", err)
	}

	isValid := Validate([]byte(randMAC))

	if !isValid {
		t.Errorf("Generated an invalid random MAC address")
	}
}