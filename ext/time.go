package ext

import (
	"time"
)

//goland:noinspection GoUnusedExportedFunction
func TimeWithOffsetFromLocal(input time.Time, location *time.Location) time.Time {
	timeLocal := input.In(location)
	_, zoneOffset := timeLocal.Zone()
	timeLocal = timeLocal.Add(time.Duration(zoneOffset) * time.Second * -1)

	return timeLocal
}

//goland:noinspection GoUnusedExportedFunction
func TimeWithoutOffsetFromUTC(input time.Time, location *time.Location) time.Time {
	timeLocal := input.In(location)
	_, zoneOffset := timeLocal.Zone()
	timeLocal = timeLocal.Add(time.Duration(zoneOffset) * time.Second)

	return timeLocal
}

// ParseDuration parsed a string formatted as a time interval, eg. 2h or 1h30m10s into a time.Duration, if the value cannot be parsed the 'def' is returned
//
//goland:noinspection GoUnusedExportedFunction
func ParseDuration(value string, def time.Duration) time.Duration {
	dur, err := time.ParseDuration(value)
	if err != nil {
		return def
	}
	return dur
}
