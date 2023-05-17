package mac

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type table struct {
		mac     string
		isValid bool
	}

	tests := []table{
		{"00-00-00-00-00-00", true},
		{"00:00:00:00:00:00", true},
		{"aa:10:89:ad:00:ff", true},
		{"FF:AB:CD:E4:80:90", true},
		{"FF-AB-CD-E4-80-90", true},
		{"fg:00:00:00:00:00", false},    // out of range
		{"ZZ-AB-CD-E4-80-90", false},    // out of range
		{"PM:AG:CD:E4:80:90", false},    // out of range
		{"aa:aa:aa:aa:aa", false},       // too short
		{"aa:aa:aa:aa:aa:aa:aa", false}, // too long
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
		mac    string
		expect string
	}

	tests := []table{
		{"aa-bc-dc-12-12-12", "AA:BC:DC:12:12:12"},
		{"aa:bc:dc:12:12:12", "AA:BC:DC:12:12:12"},
		{"AA:BC:DC:12:12:12", "AA:BC:DC:12:12:12"},
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

	isValid := Validate(randMAC)

	if !isValid {
		t.Errorf("Generated an invalid random MAC address")
	}
}
