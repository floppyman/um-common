package types

import "time"

// SqlMinDateTime defines the minimum date time allowed in MsSql
//
//goland:noinspection GoUnusedGlobalVariable
var SqlMinDateTime = time.Date(1753, 1, 1, 0, 0, 0, 0, time.UTC)

//goland:noinspection GoUnusedConst
const (
	// TimeFormatIsoDateOnly defines the date only format 2006-01-02
	TimeFormatIsoDateOnly = "2006-01-02"

	// TimeFormatIsoFullDateTimeWithDots defines the full date/time with ms format 2006-01-02T15:04:05.000
	TimeFormatIsoFullDateTimeWithDots = "2006.01.02 15:04:05.000"

	// TimeFormatIsoFullDateTimeTimezone defines the full date/time with ms and timezone format 2006-01-02T15:04:05.000-0700
	TimeFormatIsoFullDateTimeTimezone = "2006-01-02T15:04:05.000-0700"

	// TimeFormatUtcFullDateTime defines the full date/time without ms and timezone format 2006-01-02T15:04:05Z
	TimeFormatUtcFullDateTime = "2006-01-02T15:04:05Z"
)
