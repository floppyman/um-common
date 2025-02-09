package ext

import (
	"time"
)

// TimeWithOffsetFromLocal returns the 'input' time with an offset of the defined 'location'
//
//goland:noinspection GoUnusedExportedFunction
func TimeWithOffsetFromLocal(input time.Time, location *time.Location) time.Time {
	timeLocal := input.In(location)
	_, zoneOffset := timeLocal.Zone()
	timeLocal = timeLocal.Add(time.Duration(zoneOffset) * time.Second * -1)
	return timeLocal
}

// TimeWithoutOffsetFromUTC returns the 'input' time without offset from utc
//
//goland:noinspection GoUnusedExportedFunction
func TimeWithoutOffsetFromUTC(input time.Time, location *time.Location) time.Time {
	timeLocal := input.In(location)
	_, zoneOffset := timeLocal.Zone()
	timeLocal = timeLocal.Add(time.Duration(zoneOffset) * time.Second)
	return timeLocal
}

// ParseDuration takes a duration string representation and converts it into a time.Duration, or returns the 'def' if it fails
// Examples such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
//
//goland:noinspection GoUnusedExportedFunction
func ParseDuration(value string, def time.Duration) time.Duration {
	dur, err := time.ParseDuration(value)
	if err != nil {
		return def
	}
	return dur
}
