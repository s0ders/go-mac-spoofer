<img alt="go version badge" src="https://img.shields.io/github/go-mod/go-version/s0ders/go-mac-spoofer"> <img alt="go report card" src="http://goreportcard.com/badge/github.com/s0ders/go-mac-spoofer"> <img alt="license badge" src="https://img.shields.io/github/license/s0ders/go-mac-spoofer">
## Go MAC Spoofer

Open source tool written in Go to spoof a NIC MAC address on Windows, MacOS and Linux.

### Install

If Go is installed on your system, you can install this program from source as show bellow.
```bash
$ go install github.com/s0ders/go-mac-spoofer/cmd/
```

### Usage

To list available NIC on your OS, use the `list` command.
```bash
$ go-mac-spoofer list
```

To change a NIC MAC address, use the `change`command.
```bash
# change lets you specify the interface name and the new address
$ go-mac-spoofer change en0 00:1f:bd:34:10
# you can also change to a random address
$ go-mac-spoofer change -random en0
```