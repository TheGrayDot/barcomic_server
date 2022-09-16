package server

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func sendKeys(barcode string) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, must wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	for _, digit := range barcode {
		// Select keys to be pressed
		switch digit {
		case '0':
			kb.SetKeys(keybd_event.VK_0)
		case '1':
			kb.SetKeys(keybd_event.VK_1)
		case '2':
			kb.SetKeys(keybd_event.VK_2)
		case '3':
			kb.SetKeys(keybd_event.VK_3)
		case '4':
			kb.SetKeys(keybd_event.VK_4)
		case '5':
			kb.SetKeys(keybd_event.VK_5)
		case '6':
			kb.SetKeys(keybd_event.VK_6)
		case '7':
			kb.SetKeys(keybd_event.VK_7)
		case '8':
			kb.SetKeys(keybd_event.VK_8)
		case '9':
			kb.SetKeys(keybd_event.VK_9)
		}

		// Press, then release the keys
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}

	// Press return
	kb.SetKeys(keybd_event.VK_ENTER)
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}
