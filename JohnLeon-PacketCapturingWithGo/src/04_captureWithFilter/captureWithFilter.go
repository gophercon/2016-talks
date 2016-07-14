package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device       string = "en0"
	snapshot_len int32  = 65535
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = -1 * time.Second
	handle       *pcap.Handle
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	err = handle.SetBPFFilter("tcp and port 80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing TCP port 80 packets.")

	// Process packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
