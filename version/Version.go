package version

// Do not change anything in this file,
// as it is overwritten during each build.

import (
	"fmt"
	"time"
)

// Date represents the date and time the app was built.
var Date = time.Date(2017, 10, 19, 14, 25, 17, 0, time.Local)

var Major    int32 = 1
var Minor    int32 = 02
var Build    int64 = 171019
var Revision int64 = 1425

// ShortString returns the full version: 'major.minor.date.time'.
func ShortString() string {
	return fmt.Sprintf("%d.%02d", Major, Minor)
}

// FullString returns the full version: 'major.minor.date.time'.
func FullString() string {
	return fmt.Sprintf("%d.%02d.%06d.%04d", Major, Minor, Build, Revision)
}
