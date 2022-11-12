package server

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mdp/qrterminal"
)

type Config struct {
	addr        string
	port        string
	interactive bool
}

var config Config

var Version = "dev"
var Hash = "mode"

func Server() {
	// Parse the CLI arguments
	handleArguments()

	// Validate IP address and port before starting server
	if !validateAddr(config.addr) {
		fmt.Printf("[*] Error: Invalid IP address %s\n", config.addr)
		os.Exit(1)
	}
	if !validatePort(config.port) {
		fmt.Printf("[*] Error: Invalid port %s\n", config.port)
		os.Exit(1)
	}

	printQRCode(config.addr, config.port)

	// Start the HTTP API
	restAPI()
}

func handleArguments() {
	addr := flag.String("a", "0.0.0.0", "IP address to listen on")
	port := flag.String("p", "9999", "Port to listen on")
	interactive := flag.Bool("i", true, "Run interactive configuration")
	version := flag.Bool("v", false, "Prints current version")
	flag.Parse()

	if *version {
		fmt.Printf("%s-%s\n", Version, Hash)
		os.Exit(0)
	}

	config.addr = *addr
	config.port = *port
	config.interactive = *interactive

	// If requested, run the interactive server config
	if !config.interactive {
		return
	} else {
		runInteractiveConfig()
	}
}

func runInteractiveConfig() {
	// Get all network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("selectIpAddress: %+v\n", err.Error()))
		return
	}

	interfacesMap := make(map[string]string)

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
					interfacesMap[ipnet.IP.String()] = i.Name
				}
			}
		}
	}

	index := 1
	fmt.Printf("[*] The following addresses are available...\n")
	for address, networkAdapterName := range interfacesMap {
		fmt.Printf("    [%d] %s (%s)\n", index, address, networkAdapterName)
		index = index + 1
	}

	fmt.Print("[*] Enter IP address: ")
	var addr string
	fmt.Scanln(&addr)
	config.addr = strings.Trim(addr, " ")

	fmt.Print("[*] Enter port: ")
	var port string
	fmt.Scanln(&port)
	config.port = strings.Trim(port, " ")
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

func printQRCode(addr, port string) {
	qrconfig := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	host := addr + ":" + port
	qrterminal.GenerateWithConfig(host, qrconfig)
}
