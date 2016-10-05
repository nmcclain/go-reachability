package reachability

import (
	"fmt"
	"net"
	"time"
)

func IsReachable(host string, port int) error {
	return IsReachableTimeout(host, port, time.Second*10)
}

func IsReachableTimeout(host string, port int, timeout time.Duration) error {
	if len(host) < 1 {
		return fmt.Errorf("must specify a host")
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		return fmt.Errorf("TCP connection error: %s", err.Error())
	}
	defer func() {
		conn.Close()
	}()
	return nil
}
