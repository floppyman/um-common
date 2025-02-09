package sqls

import (
	"database/sql"
	"time"
)

// EmptyNullTime creates a new sql.NullTime that is DbNull
//
//goland:noinspection GoUnusedExportedFunction
func EmptyNullTime() sql.NullTime {
	return sql.NullTime{Valid: false}
}

// EmptyNullFloat64 creates a new sql.NullFloat64 that is DbNull
//
//goland:noinspection GoUnusedExportedFunction
func EmptyNullFloat64() sql.NullFloat64 {
	return sql.NullFloat64{Valid: false}
}

// EmptyNullString creates a new sql.NullString that is DbNull
//
//goland:noinspection GoUnusedExportedFunction
func EmptyNullString() sql.NullString {
	return sql.NullString{Valid: false}
}

// EmptyNullInt32 creates a new sql.NullInt32 that is DbNull
//
//goland:noinspection GoUnusedExportedFunction
func EmptyNullInt32() sql.NullInt32 {
	return sql.NullInt32{Valid: false}
}

// EmptyNullInt64 creates a new sql.NullInt64 that is DbNull
//
//goland:noinspection GoUnusedExportedFunction
func EmptyNullInt64() sql.NullInt64 {
	return sql.NullInt64{Valid: false}
}

// EmptyNullBool creates a new sql.NullBool that is DbNull
//
//goland:noinspection GoUnusedExportedFunction
func EmptyNullBool() sql.NullBool {
	return sql.NullBool{Valid: false}
}

// NewNullTime creates a new sql.NullTime that is not DbNull and contains the 'value'
//
//goland:noinspection GoUnusedExportedFunction
func NewNullTime(value time.Time) sql.NullTime {
	return sql.NullTime{Time: value, Valid: true}
}

// NewNullFloat64 creates a new sql.NullFloat64 that is not DbNull and contains the 'value'
//
//goland:noinspection GoUnusedExportedFunction
func NewNullFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}

// NewNullString creates a new sql.NullString that is not DbNull and contains the 'value'
//
//goland:noinspection GoUnusedExportedFunction
func NewNullString(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}

// NewNullInt32 creates a new sql.NullInt32 that is not DbNull and contains the 'value'
//
//goland:noinspection GoUnusedExportedFunction
func NewNullInt32(value int32) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: true}
}

// NewNullInt64 creates a new sql.NullInt64 that is not DbNull and contains the 'value'
//
//goland:noinspection GoUnusedExportedFunction
func NewNullInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: true}
}

// NewNullBool creates a new sql.NullBool that is not DbNull and contains the 'value'
//
//goland:noinspection GoUnusedExportedFunction
func NewNullBool(value bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: true}
}

// NewNullTimeWithValid creates a new sql.NullTime with the defined 'value' and 'valid' values
//
//goland:noinspection GoUnusedExportedFunction
func NewNullTimeWithValid(value time.Time, valid bool) sql.NullTime {
	return sql.NullTime{Time: value, Valid: valid}
}

// NewNullFloat64WithValid creates a new sql.NullFloat64 with the defined 'value' and 'valid' values
//
//goland:noinspection GoUnusedExportedFunction
func NewNullFloat64WithValid(value float64, valid bool) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: valid}
}

// NewNullStringWithValid creates a new sql.NullString with the defined 'value' and 'valid' values
//
//goland:noinspection GoUnusedExportedFunction
func NewNullStringWithValid(value string, valid bool) sql.NullString {
	return sql.NullString{String: value, Valid: valid}
}

// NewNullInt32WithValid creates a new sql.NullInt32 with the defined 'value' and 'valid' values
//
//goland:noinspection GoUnusedExportedFunction
func NewNullInt32WithValid(value int32, valid bool) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: valid}
}

// NewNullInt64WithValid creates a new sql.NullInt64 with the defined 'value' and 'valid' values
//
//goland:noinspection GoUnusedExportedFunction
func NewNullInt64WithValid(value int64, valid bool) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: valid}
}

// NewNullBoolWithValid creates a new sql.NullBool with the defined 'value' and 'valid' values
//
//goland:noinspection GoUnusedExportedFunction
func NewNullBoolWithValid(value bool, valid bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: valid}
}

// ValueOrDefaultNullTime checks if the sql.NullTime is valid and returns its value, else returning the default sql.NullTime
//
//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullTime(value sql.NullTime) time.Time {
	if value.Valid {
		return value.Time
	}
	return time.Time{}
}

// ValueOrDefaultNullFloat64 checks if the sql.NullFloat64 is valid and returns its value, else returning the default sql.NullFloat64
//
//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullFloat64(value sql.NullFloat64) float64 {
	if value.Valid {
		return value.Float64
	}
	return 0
}

// ValueOrDefaultNullString checks if the sql.NullString is valid and returns its value, else returning the default sql.NullString
//
//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

// ValueOrDefaultNullInt32 checks if the sql.NullInt32 is valid and returns its value, else returning the default sql.NullInt32
//
//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullInt32(value sql.NullInt32) int32 {
	if value.Valid {
		return value.Int32
	}
	return 0
}

// ValueOrDefaultNullInt64 checks if the sql.NullInt64 is valid and returns its value, else returning the default sql.NullInt64
//
//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullInt64(value sql.NullInt64) int64 {
	if value.Valid {
		return value.Int64
	}
	return 0
}

// ValueOrDefaultNullBool checks if the sql.NullBool is valid and returns its value, else returning the default sql.NullBool
//
//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullBool(value sql.NullBool) bool {
	if value.Valid {
		return value.Bool
	}
	return false
}
