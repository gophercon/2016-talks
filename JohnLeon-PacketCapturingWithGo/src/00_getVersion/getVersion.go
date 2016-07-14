package main

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func main() {
	version := pcap.Version()
	fmt.Println(version)
}
