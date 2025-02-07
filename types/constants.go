package types

import "time"

//goland:noinspection GoUnusedGlobalVariable
var SqlMinDateTime = time.Date(1753, 1, 1, 0, 0, 0, 0, time.UTC)

//goland:noinspection GoUnusedGlobalVariable
const (
	TimeFormatIsoDateOnly             = "2006-01-02"
	TimeFormatIsoFullDateTimeTimezone = "2006-01-02T15:04:05.000-0700"
	TimeFormatIsoFullDateTimeWithDots = "2006.01.02 15:04:05.000"
)
