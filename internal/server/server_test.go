package server

import (
	"testing"
)

func TestValidateAddr(t *testing.T) {
	table := []struct {
		addr     string
		expected bool
	}{
		{"192.168.1.100", true},
		{"0.0.0.0", true},
		{"204.212.2121.12", false},
	}

	for _, v := range table {
		t.Run(v.addr, func(t *testing.T) {
			result := validateAddr(v.addr)
			if result != v.expected {
				t.Fatalf("Expected result: %t, but got: %t", v.expected, result)
			}
		})
	}
}

func TestValidatePort(t *testing.T) {
	table := []struct {
		port     string
		expected bool
	}{
		{"9999", true},
		{"100", true},
		{"64000", true},
		{"-1", false},
		{"65999", false},
	}

	for _, v := range table {
		t.Run(v.port, func(t *testing.T) {
			result := validatePort(v.port)
			if result != v.expected {
				t.Fatalf("Expected result: %t, but got: %t", v.expected, result)
			}
		})
	}
}
