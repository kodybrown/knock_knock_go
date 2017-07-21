package version

// Do not change anything in this file,
// as it is overwritten during each build.

import (
	"fmt"
	"time"
)

// Date represents the date and time the app was built.
var Date = time.Date(2017, 7, 21, 11, 31, 45, 0, time.Local)

var Major    int32 = 0
var Minor    int32 = 7
var Build    int64 = 2017721
var Revision int64 = 1131

// ShortString returns the full version: 'major.minor.date.time'.
func ShortString() string {
	return fmt.Sprintf("02d", Major, Minor)
}

// FullString returns the full version: 'major.minor.date.time'.
func FullString() string {
	return fmt.Sprintf("02d.:update_version6d.:update_version4d", Major, Minor, Build, Revision)
}
