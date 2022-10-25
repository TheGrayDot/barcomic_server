package server

import (
	"testing"
)

func TestInitalizeKeys(t *testing.T) {
	keyBonding = InitalizeKeys()
	SendKeys("", keyBonding)
}
