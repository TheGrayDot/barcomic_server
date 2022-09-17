package server

import (
	"fmt"
)

var ConnHost string = "0.0.0.0"
var ConnPort string = "9999"

func Server() {
	interactiveInput()
	server()
}

func interactiveInput() {
	fmt.Print("[*] Start server using default 0.0.0.0:9999 [Y/n]")
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
