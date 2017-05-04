/*
    Copyright (C) 2017 Kody Brown

    Released under the MIT License:

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
*/

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

//
// knock_knock.exe hostname port1 port2 ...
//

func main() {
	host := os.Args[1]
	ports := []int{}

	fmt.Println("Host:")
	fmt.Printf("  %s\n", host)

	fmt.Println("Ports:")
	for _, arg := range os.Args[2:] {
		fmt.Printf("  %s\n", arg)
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(11)
		}
		ports = append(ports, i)
	}
	fmt.Println("")

	for _, i := range ports {
		connectTo(host, i)
	}

	os.Exit(0)
}

func connectTo(host string, port int) {
	fmt.Printf("knocking on %s:%d..\n", host, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Millisecond * 10)
	if err != nil {
		// ignore error
		return
	}
	defer conn.Close()
}
