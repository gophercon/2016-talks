package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device       string = "en0"
	snapshot_len int32  = 65535
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = -1 * time.Second // negative will return immediately instead of waiting
	handle       *pcap.Handle
)

func main() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {

		fmt.Println("====== Packet Layers =======")
		for _, layer := range packet.Layers() {
			fmt.Println(layer.LayerType())
		}
		fmt.Println("============================")

		// Get IPv4 layer
		ip4Layer := packet.Layer(layers.LayerTypeIPv4)
		if ip4Layer != nil {
			fmt.Println("IPv4 layer detected.")
			ip, _ := ip4Layer.(*layers.IPv4)
			fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
			fmt.Println("Protocol: ", ip.Protocol)
			fmt.Println()
		} else {
			fmt.Println("No IPv4 layer detected.")
		}

		// Get TCP layer
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer != nil {
			fmt.Println("TCP layer detected.")
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Println("ACK: ", tcp.ACK)
			fmt.Println("SYN: ", tcp.SYN)
			fmt.Println("Seq: ", tcp.Seq)
			fmt.Println("DstPort: ", tcp.DstPort)
			fmt.Println("SrcPort: ", tcp.SrcPort)
			fmt.Println()
		} else {
			fmt.Println("No TCP layer detected.")
		}
	}
}
