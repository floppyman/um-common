package sqls

import (
	"database/sql"
	"time"
)

//goland:noinspection GoUnusedExportedFunction
func EmptyNullTime() sql.NullTime {
	return sql.NullTime{Valid: false}
}

//goland:noinspection GoUnusedExportedFunction
func EmptyNullFloat64() sql.NullFloat64 {
	return sql.NullFloat64{Valid: false}
}

//goland:noinspection GoUnusedExportedFunction
func EmptyNullString() sql.NullString {
	return sql.NullString{Valid: false}
}

//goland:noinspection GoUnusedExportedFunction
func EmptyNullInt32() sql.NullInt32 {
	return sql.NullInt32{Valid: false}
}

//goland:noinspection GoUnusedExportedFunction
func EmptyNullInt64() sql.NullInt64 {
	return sql.NullInt64{Valid: false}
}

//goland:noinspection GoUnusedExportedFunction
func EmptyNullBool() sql.NullBool {
	return sql.NullBool{Valid: false}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullTime(value time.Time) sql.NullTime {
	return sql.NullTime{Time: value, Valid: true}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullString(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullInt32(value int32) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: true}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: true}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullBool(value bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: true}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullTimeWithValid(value time.Time, valid bool) sql.NullTime {
	return sql.NullTime{Time: value, Valid: valid}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullFloat64WithValid(value float64, valid bool) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: valid}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullStringWithValid(value string, valid bool) sql.NullString {
	return sql.NullString{String: value, Valid: valid}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullInt32WithValid(value int32, valid bool) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: valid}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullInt64WithValid(value int64, valid bool) sql.NullInt64 {
	return sql.NullInt64{Int64: value, Valid: valid}
}

//goland:noinspection GoUnusedExportedFunction
func NewNullBoolWithValid(value bool, valid bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: valid}
}

//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullTime(value sql.NullTime) time.Time {
	if value.Valid {
		return value.Time
	}
	return time.Time{}
}

//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullFloat64(value sql.NullFloat64) float64 {
	if value.Valid {
		return value.Float64
	}
	return 0
}

//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullInt32(value sql.NullInt32) int32 {
	if value.Valid {
		return value.Int32
	}
	return 0
}

//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullInt64(value sql.NullInt64) int64 {
	if value.Valid {
		return value.Int64
	}
	return 0
}

//goland:noinspection GoUnusedExportedFunction
func ValueOrDefaultNullBool(value sql.NullBool) bool {
	if value.Valid {
		return value.Bool
	}
	return false
}
