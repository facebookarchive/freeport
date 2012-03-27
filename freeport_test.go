package freeport

import (
	"testing"
	"net"
	"strconv"
)

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		t.Fatalf("Got err from GetFreePort: %s", err)
	}

	if port == 0 {
		t.Fatal("Got port 0")
	}
}

func TestGetFreePortIsClosed(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		t.Fatalf("Got err from GetFreePort: %s", err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:" + strconv.Itoa(port))
  if err != nil {
		t.Fatalf("Got err from net.Listen: %s", err)
	}
	listener.Close()
}
