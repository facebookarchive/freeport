// This library utilizes the feature of most OSs where a request to
// bind to port 0 will trigger the kernel to provide a free port.
package freeport

import (
	"net"
	"strconv"
)

// Get a free port.
func Get() (port int, err error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().String()
	_, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(portString)
}
