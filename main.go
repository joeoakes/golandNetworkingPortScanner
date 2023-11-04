package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(protocol, hostname string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		fmt.Printf("Port %d is closed or filtered.\n", port)
		return
	}
	conn.Close()
	fmt.Printf("Port %d is open.\n", port)
}

func main() {
	var wg sync.WaitGroup
	protocol := "tcp"
	hostname := "scanme.nmap.org"

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go scanPort(protocol, hostname, port, &wg)
	}
	wg.Wait()
}
