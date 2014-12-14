//
// Copyright 2014 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package serial

import "go.bug.st/serial"
import "fmt"
import "log"

func ExampleCommunication() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	port, err := serial.OpenPort(ports[0], false)
	if err != nil {
		log.Fatal(err)
	}
	if err = port.SetSpeed(9600); err != nil {
		log.Fatal(err)
	}
	if err := port.SetParity(serial.PARITY_NONE); err != nil {
		log.Fatal(err)
	}
	if err := port.SetDataBits(8); err != nil {
		log.Fatal(err)
	}
	if err := port.SetStopBits(serial.STOPBITS_ONE); err != nil {
		log.Fatal(err)
	}
	n, err := port.Write([]byte("10,20,30\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
	}
}

// vi:ts=2
