package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/TheGrayDot/barcomic/internal/barcomic"
)

var Version = "dev"
var Hash = "mode"

func main() {
	fmt.Printf("[*] barcomic %s-%s\n", Version, Hash)

	// Configure command line arguments
	addr := flag.String("a", "", "Address to listen on")
	port := flag.String("p", "9999", "Port to listen on")
	disableKeystrokes := flag.Bool("k", false, "Disable keystrokes")
	interactive := flag.Bool("i", true, "Run interactive configuration")
	verbose := flag.Bool("v", false, "Prints verbose information")
	flag.Parse()

	// If address is provided, set interactive to false
	if *addr != "" {
		*interactive = false
	}

	// If requested, run the interactive server config
	if *interactive {
		*addr = interactiveNetworkConfiguration(*addr)
	}

	// Validate IP address and port before starting server
	if !validateAddr(*addr) {
		fmt.Printf("[*] Error: Invalid IP address %s\n", *addr)
		os.Exit(1)
	}
	if !validatePort(*port) {
		fmt.Printf("[*] Error: Invalid port %s\n", *port)
		os.Exit(1)
	}

	barcomic.Start(*addr, *port, *disableKeystrokes, *verbose)
}

func interactiveNetworkConfiguration(addr string) string {
	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("selectIpAddress: %+v\n", err.Error()))
		return "0.0.0.0"
	}

	var availableInterfaces [][2]string

	// Loop through interfaces
	for _, i := range interfaces {
		// Get interface name (e.g., wlan0)
		byNameInterface, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}

		// Get addresses on interface and loop
		addresses, err := byNameInterface.Addrs()
		for _, address := range addresses {
			// Check address is IP4 and not loopback (127.0.0.1)
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					var availableInterface [2]string
					availableInterface[0] = ipnet.IP.String()
					availableInterface[1] = i.Name
					availableInterfaces = append(availableInterfaces, availableInterface)
				}
			}
		}
	}

	fmt.Printf("[*] The following addresses are available...\n")
	for i, availableInterface := range availableInterfaces {
		fmt.Printf("    [%d] %s (%s)\n", i, availableInterface[0], availableInterface[1])
	}

	// Get address selection from user
	fmt.Print("[*] Enter IP address [0.0.0.0]: ")
	var addrInput string
	fmt.Scanln(&addrInput)
	addr = strings.Trim(addrInput, " ")

	// Check for valid input address
	validAddress := validateAddr(addrInput)
	if validAddress {
		return addrInput
	}

	addrInputInt, err := strconv.Atoi(addrInput)
	if err != nil {
		fmt.Println("\n[*] Error. Could not determine network interface selection.")
		os.Exit(1)
	}
	if addrInputInt < 0 || addrInputInt >= len(availableInterfaces) {
		fmt.Println("\n[*] Error: Network interface selection error.")
		os.Exit(1)
	}
	addrInput = availableInterfaces[addrInputInt][0]
	return addrInput
}

func validateAddr(addr string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(addr) {
		return true
	}
	return false
}

func validatePort(port string) bool {
	portInt, err := strconv.ParseInt(port, 10, 0)
	if err != nil {
		return false
	}
	if portInt >= 0 && portInt <= 65535 {
		return true
	}
	return false
}
