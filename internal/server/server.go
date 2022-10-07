package server

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	addr        string
	port        string
	interactive bool
}

var config Config

func Server() {
	// Parse the CLI arguments
	handleArguments()

	// If requested, run the interactive server config
	if config.interactive {
		runInteractiveConfig()
	}

	// Validate IP address and port before starting server
	if !validateAddr(config.addr) {
		fmt.Printf("[*] Error: Invalid IP address %s\n", config.addr)
		os.Exit(1)
	}
	if !validatePort(config.port) {
		fmt.Printf("[*] Error: Invalid port %s\n", config.port)
		os.Exit(1)
	}

	// Start the HTTP API
	restAPI()
}

func handleArguments() {
	addr := flag.String("a", "0.0.0.0", "IP address to listen on")
	port := flag.String("p", "9999", "Port to listen on")
	interactive := flag.Bool("i", false, "Run interactive configuration")
	flag.Parse()
	config.addr = *addr
	config.port = *port
	config.interactive = *interactive
}

func runInteractiveConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("[*] Start server using default %s:%s [Y/n]", config.addr, config.port)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	response = strings.ToLower(response)

	if response == "y" || response == "" {
		return
	} else {
		fmt.Print("[*] Enter IP address: ")
		var addr string
		fmt.Scanln(&addr)
		config.addr = strings.Trim(addr, " ")

		fmt.Print("[*] Enter port: ")
		var port string
		fmt.Scanln(&port)
		config.port = strings.Trim(port, " ")
	}
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
