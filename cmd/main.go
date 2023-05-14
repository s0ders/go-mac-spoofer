package main

import (
	"fmt"
	"os"

	"github.com/s0ders/go-mac-spoofer/internal/nic"
)

func main() {

	args := os.Args[1:]
	cmd := args[0]


	switch(cmd) {
	case "list":
		nics, err := nic.ListNICs()

		if err != nil {
			panic(err)
		}

		for i, nic := range nics {
			fmt.Printf("(%d) %s\n", i+1, string(nic))
		}
	}

}
