package main

import (
	"fmt"
	serial "go.bug.st/serial.v1"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s SERIAL_PORT\n", os.Args[0])
		return
	}
	port := os.Args[1]
	p, err := serial.Open(port, &serial.Mode{BaudRate: 1200})
	if err != nil {
		panic("Error opening port")
	}
	defer p.Close()

	if err = p.SetDTR(false); err != nil {
		panic("cannot set DTR")
	}
}
