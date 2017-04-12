package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dmichael/go-multicast/multicast"
	"github.com/urfave/cli"
)

const (
	defaultMulticastAddress = "239.0.0.0:9999"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)
		if address == "" {
			address = defaultMulticastAddress
		}
		fmt.Printf("Broadcasting to %s\n", address)
		ping(address)
		return nil
	}

	app.Run(os.Args)
}

func ping(addr string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn.Write([]byte("hello, world\n"))
		time.Sleep(1 * time.Second)
	}
}
