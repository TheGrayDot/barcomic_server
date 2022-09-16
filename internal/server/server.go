package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var connHost string = "0.0.0.0"
var connPort string = "9999"

func Server() {
	interactiveInput()
	startServer()
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
		connHost = tempHost
		connPort = tempPort
	}
}

func startServer() {
	fmt.Println("[*] Starting server on " + connHost + ":" + connPort)

	l, err := net.Listen("tcp", connHost+":"+connPort)
	if err != nil {
		fmt.Println("[*] Error listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("[*] Error connecting:", err.Error())
			return
		}
		fmt.Println("[*] Client connected...")
		fmt.Println("[*] Client information: " + c.RemoteAddr().String())

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("[*] Client left...")
		conn.Close()
		return
	}

	fmt.Println("[*] Client message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	bufferString := string(string(buffer[:len(buffer)-1]))
	sendKeys(bufferString)

	handleConnection(conn)
}
