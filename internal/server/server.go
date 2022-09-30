package server

import (
	"fmt"
)

// Set server defaults
var ConnHost string = "0.0.0.0"
var ConnPort string = "9999"

func Server() {
	interactiveInput()
	restapi()
}

func interactiveInput() {
	fmt.Printf("[*] Start server using default %s:%s [Y/n]", ConnHost, ConnPort)
	var defaults string
	fmt.Scanln(&defaults)
	if defaults == "" || defaults == "y" {
		defaults = "Y"
	}
	if defaults != "Y" {
		var tempHost string
		fmt.Print("[*] Enter IP address: ")
		fmt.Scanln(&tempHost)
		var tempPort string
		fmt.Print("[*] Enter port: ")
		fmt.Scanln(&tempPort)
		ConnHost = tempHost
		ConnPort = tempPort
	}
}
