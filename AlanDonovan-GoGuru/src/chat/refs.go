package main

import network "net"

func _() {
	network.Listen("tcp", "localhost:8000")
}
