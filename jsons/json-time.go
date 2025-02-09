package jsons

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime struct {
	Value time.Time `json:"value"`
	Valid bool      `json:"valid"`
	Set   bool      `json:"set"`
}

// NewJsonTime creates a new instance of JsonTime with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonTime(value time.Time) *JsonTime {
	return &JsonTime{
		Value: time.Date(value.Year(), value.Month(), value.Day(), value.Hour(), value.Minute(), value.Second(), value.Nanosecond(), value.Location()),
		Valid: true,
		Set:   true,
	}
}

// NullJsonTime creates a new instance of JsonTime that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonTime() *JsonTime {
	return &JsonTime{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonTime to Json
func (i *JsonTime) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", i.Value.Format(time.RFC3339Nano))), nil
}

// UnmarshalJSON converts from Json to JsonTime
func (i *JsonTime) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp time.Time
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of time.Time 'time.Time{}'
func (i *JsonTime) ValueOrDefault() time.Time {
	return i.ValueOrDefaultValue(time.Time{}) // time.Time{} = minimum time 1 jan 0001 00:00:00 UTC
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonTime) ValueOrDefaultValue(val time.Time) time.Time {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonTime is valid and set
func (i *JsonTime) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonTime) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return i.Value.Format(time.RFC3339)
}
