package sysctl

import (
	"fmt"
	"runtime"
	"testing"
)

func TestValueFromKey(t *testing.T) {
	var param string
	switch runtime.GOOS {
	case `darwin`:
		param = `net.inet.ip.forwarding`
	case `linux`:
		param = `net.ipv4.ip_forward`
	}
	value, err := Get(param)
	if err != nil {
		t.Fatalf("Error Received: %v", err)
	}
	fmt.Println(">>>", param, ">", value)
}
func TestPathFromKey(t *testing.T) {
	in := "net.ipv4.ip_forward"
	expected := "/proc/sys/net/ipv4/ip_forward"
	got := pathFromKey(in)
	if got != expected {
		t.Fatalf("Expected: %s. Got: %s", expected, got)
	}
}

func TestKeyFromPath(t *testing.T) {
	in := "/proc/sys/net/ipv4/ip_forward"
	expected := "net.ipv4.ip_forward"
	got := keyFromPath(in)
	if got != expected {
		t.Fatalf("Expected: %s. Got: %s", expected, got)
	}
}
