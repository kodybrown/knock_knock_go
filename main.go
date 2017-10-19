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
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//
// knock.exe hostname port1 port2 ...
//
// or
//
// knock.exe -i [hostname]
//

func init() {
	ParseFlags()
}

func main() {
	var host string
	var ports = []int{}

	var err error

	if optInteractive {
		host = ""

		if len(os.Args) > 2 {
			host = os.Args[2]
		}
		if host == "" {
			host, err = getInput("Enter the hostname: ")
			if err != nil {
				os.Exit(11)
			}
		}

		ports, err = getPorts()
		if err != nil {
			os.Exit(11)
		}
		fmt.Println("")
	} else {
		host = os.Args[1]

		for _, arg := range os.Args[2:] {
			i, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(11)
			}
			ports = append(ports, i)
		}
	}

	if host == "" || len(ports) == 0 {
		fmt.Println("A hostname and port is required.")
		os.Exit(11)
	}

	fmt.Printf("Host:\n  %s\n", host)
	fmt.Printf("Ports:\n  %v\n", ports)
	fmt.Println("")

	if optInteractive {
		result, err := getInput("Continue [Y/n]: ")
		if err != nil {
			os.Exit(12)
		}
		if result != "y" && result != "Y" && result != "" {
			os.Exit(0)
		}
	}

	fmt.Println("")
	for _, p := range ports {
		connectTo(host, p)
	}

	os.Exit(0)
}

func getInput(prompt string) (input string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	input, err = reader.ReadString('\n')
	if err != nil {
		return input, err
	}

	input = strings.Trim(input, " \r\n\t")
	return input, nil
}

func getPorts() (ports []int, err error) {
	ports = []int{}
	err = nil

	for true {
		tmp, err := getInput("Enter port: ")
		if err != nil {
			return ports, err
		}
		tmp = strings.Trim(tmp, " \r\n\t")
		if tmp == "" || tmp == "-1" {
			break
		}
		port, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Printf("INVALID port specified: `%s`\n", tmp)
			return ports, err
		}
		ports = append(ports, port)
	}

	return ports, nil
}

func connectTo(host string, port int) {
	fmt.Printf("knocking on %s:%d..\n", host, port)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Millisecond*10)
	if err != nil {
		// ignore error
		return
	}
	defer conn.Close()
}
