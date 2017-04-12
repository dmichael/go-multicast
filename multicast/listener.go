// https://gist.github.com/fiorix/9664255
// https://en.wikipedia.org/wiki/Multicast_address
// https://support.mcommstv.com/hc/en-us/articles/202306226-Choosing-multicast-addresses-and-ports

package multicast

import (
	"log"
	"net"
)

const (
	maxDatagramSize = 8192
)

// Listen binds to the UDP address and port given and writes packets received
// from that address to a buffer which is passed to a hander
func Listen(address string, handler func(*net.UDPAddr, int, []byte)) {
	// go ping(address)
	// serveMulticastUDP(address, msgHandler)
	// Parse the string address
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	conn.SetReadBuffer(maxDatagramSize)

	// Loop forever reading from the socket
	for {
		buffer := make([]byte, maxDatagramSize)
		numBytes, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		handler(src, numBytes, buffer)
	}
}

// func ping(a string) {
// 	addr, err := net.ResolveUDPAddr("udp", a)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c, err := net.DialUDP("udp", nil, addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		c.Write([]byte("hello, world\n"))
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func serveMulticastUDP(a string, h func(*net.UDPAddr, int, []byte)) {
// 	addr, err := net.ResolveUDPAddr("udp", a)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	l, err := net.ListenMulticastUDP("udp", nil, addr)
// 	l.SetReadBuffer(maxDatagramSize)
// 	for {
// 		b := make([]byte, maxDatagramSize)
// 		n, src, err := l.ReadFromUDP(b)
// 		if err != nil {
// 			log.Fatal("ReadFromUDP failed:", err)
// 		}
// 		h(src, n, b)
// 	}
// }
