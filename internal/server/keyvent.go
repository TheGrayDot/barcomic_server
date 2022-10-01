package server

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func InitalizeKeys() keybd_event.KeyBonding {
	// Initialize keyboard event
	keyBonding, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, must wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	return keyBonding
}

func SendKeys(barcode string, keyBonding keybd_event.KeyBonding) {
	keyBonding.Clear()
	for _, digit := range barcode {
		// Select keys to be pressed
		switch digit {
		case '0':
			keyBonding.SetKeys(keybd_event.VK_0)
		case '1':
			keyBonding.SetKeys(keybd_event.VK_1)
		case '2':
			keyBonding.SetKeys(keybd_event.VK_2)
		case '3':
			keyBonding.SetKeys(keybd_event.VK_3)
		case '4':
			keyBonding.SetKeys(keybd_event.VK_4)
		case '5':
			keyBonding.SetKeys(keybd_event.VK_5)
		case '6':
			keyBonding.SetKeys(keybd_event.VK_6)
		case '7':
			keyBonding.SetKeys(keybd_event.VK_7)
		case '8':
			keyBonding.SetKeys(keybd_event.VK_8)
		case '9':
			keyBonding.SetKeys(keybd_event.VK_9)
		}

		// Press, then release the keys
		keyBonding.Press()
		time.Sleep(10 * time.Millisecond)
		keyBonding.Release()
	}

	// Press return
	keyBonding.SetKeys(keybd_event.VK_ENTER)
	keyBonding.Press()
	time.Sleep(10 * time.Millisecond)
	keyBonding.Release()
}
