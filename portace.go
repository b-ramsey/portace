package main

import (
	"fmt"
	"flag"
	"os"
	"net"
	"time"
)

func main() {
	// Command line arguments
	addressPtr := flag.String("address", "", "IP address or hostname to scan")
	minPortPtr := flag.Int("min-port", 0, "Minimum port to scan in range")
	maxPortPtr := flag.Int("max-port", 0, "Maximum port to scan in range")

	flag.Parse()

	// Check that an address has been provided
	if *addressPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Get the ports to be scanned
	var portList = generatePortList(*minPortPtr, *maxPortPtr)

	// Scan each of the ports provided
	// TODO make this concurrent
	for _, currPort := range portList {
		testOpen(*addressPtr, currPort)
	}
}

func generatePortList(min int, max int) []int {
	// No ports specified, use the top 10
	if min == 0 || max == 0 {
		port_list := []int {10, 21, 22, 23, 25, 80, 110, 139, 443, 445, 3389}
		return port_list
	}
	// Create an empty list
	port_list := make([]int, 0)

	// Build the port range
	for port := min; port <= max; port++ {
		port_list = append(port_list, port)
	}

	return port_list
}

func testOpen(address string, port int) {
	addr := fmt.Sprintf("%s:%d", address, port)
	// Attempt to connect to the full address
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)

	// An error occurred, the port is closed
	if err != nil {
		fmt.Printf("%d:\tCLOSED\n", port)
	}
	// A connection needs to be formed to be closed, this means it is open
	if conn != nil {
		defer conn.Close()
		fmt.Printf("%d:\tOPEN\n", port)
	}

}