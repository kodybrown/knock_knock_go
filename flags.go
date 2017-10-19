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
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kodybrown/knock_knock_go/version"
)

var (
	// AppName ..
	AppName        = "knock"
	yearAppCreated = 2017

	// Environment variables NOT allowed.
	optHelp        bool
	optVersion     bool
	optVersionFull bool
	optEnvars      bool
	optDebug       bool
	optInteractive bool

	// Environment variables are allowed.

	// Ignore certain flags for some events.
	ignoreEnvFlagList  = []string{"h", "help", "v", "version", "envars", "debug", "test", "rotate"}
	ignoreHelpFlagList = []string{}
)

// ParseFlags parses the command-line flags and options
func ParseFlags() {
	// Specify the command-line arguments/flags.
	flag.BoolVar(&optHelp, "h", false, "")
	flag.BoolVar(&optHelp, "help", false, "displays this help")
	flag.BoolVar(&optVersion, "v", false, "display build version")
	flag.BoolVar(&optVersionFull, "version", false, "display build build version and copyright info")
	flag.BoolVar(&optEnvars, "envars", false, "display supported environment variables")
	flag.BoolVar(&optDebug, "debug", false, "enable debug mode")
	flag.BoolVar(&optInteractive, "i", false, "")
	flag.BoolVar(&optInteractive, "interactive", false, "interactive mode; will prompt for the host and ports")

	// Always load the envars before parsing the command-line
	// arguments/flags, so that the flags will overwrite them.
	LoadEnvars()

	// Parse the command-line arguments/flags.
	flag.Parse()

	if optHelp {
		printHeader()
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Printf("     %s [flags]\n", AppName)
		fmt.Println("")
		fmt.Println(" flags:")
		flag.VisitAll(func(f *flag.Flag) {
			if !StringInSlice(f.Name, ignoreHelpFlagList) {
				if len(f.Name) == 1 {
					fmt.Printf("    -%-16s %s\n", f.Name, f.Usage)
				} else {
					fmt.Printf("   --%-16s %s\n", f.Name, f.Usage)
				}
			}
		})
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println("     knock host port [port...]   knocks host on specified port(s)")
		fmt.Println("     knock -i [host]             launches interactive mode")
		os.Exit(0)
	} else if optEnvars {
		printHeader()
		fmt.Println("")
		fmt.Println("ENVARS:")
		fmt.Println("")
		flag.VisitAll(func(f *flag.Flag) {
			if !StringInSlice(f.Name, ignoreEnvFlagList) && !StringInSlice(f.Name, ignoreHelpFlagList) {
				envName := fmt.Sprintf("%s_%s", AppName, strings.ToUpper(f.Name))
				flagValue := os.Getenv(envName)
				if flagValue != "" {
					fmt.Printf("   %-24s %-12s %s\n", envName, flagValue, f.Usage)
				} else {
					fmt.Printf("   %-24s %-12s %s\n", envName, "(not set)", f.Usage)
				}
			}
		})
		os.Exit(0)
	} else if optVersion {
		fmt.Printf("%s\n", version.ShortString())
		os.Exit(0)
	} else if optVersionFull {
		printHeader()
		os.Exit(0)
	}

	printHeader()
}

func printHeader() {
	fmt.Printf("knock_knock_go v%s (%d.%d)\n", version.ShortString(), version.Build, version.Revision)
	if version.Date.Year() == yearAppCreated {
		fmt.Printf("Copyright (C) %d Kody Brown. All Rights Reserved.\n", version.Date.Year())
	} else {
		fmt.Printf("Copyright (C) %d-%d Kody Brown. All Rights Reserved.\n", yearAppCreated, version.Date.Year())
	}
	fmt.Println("Source: https://github.com/kodybrown/knock_knock_go. Released under the MIT license.")
}

// LoadEnvars sets flag values from the environment variable (if exists).
// This happens before the flags are processed, so command-line arguments
// will ALWAYS overwrite environment variables.
func LoadEnvars() {
	flag.VisitAll(func(f *flag.Flag) {
		if !StringInSlice(f.Name, ignoreEnvFlagList) {
			envName := fmt.Sprintf("%s_%s", strings.ToUpper(AppName), strings.ToUpper(f.Name))
			flagValue := os.Getenv(envName)
			if flagValue != "" {
				// fmt.Printf("FOUND ENVAR: %s=%s\n", envName, flagValue)
				if err := flag.Set(f.Name, flagValue); err != nil {
					fmt.Printf("**** ERROR: Could not set flag %q from environment (value: %q).\n%v\n", f.Name, flagValue, err)
				}
			}
		}
	})
}

// StringInSlice returns whether or not `a`
// is in `list`.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
