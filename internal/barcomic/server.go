package barcomic

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal"
)

type Config struct {
	addr              string
	port              string
	disableKeystrokes bool
	verbose           bool
}

var config Config

func Start(addr, port string, disableKeystrokes, verbose bool) {
	config.addr = addr
	config.port = port
	config.disableKeystrokes = disableKeystrokes
	config.verbose = verbose

	// Print QR code for user to scan
	printQRCode(config.addr, config.port)

	// Start the HTTP API
	startRestApi()
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
	fmt.Printf("[*] Starting server using %s:%s\n", config.addr, config.port)
}
