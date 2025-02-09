package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/umbrella-sh/um-common/ext"
)

type JsonFloat64 struct {
	Value float64 `json:"value"`
	Valid bool    `json:"valid"`
	Set   bool    `json:"set"`
}

// NewJsonFloat64 creates a new instance of JsonFloat64 with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonFloat64(value float64) *JsonFloat64 {
	return &JsonFloat64{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

// NullJsonFloat64 creates a new instance of JsonFloat64 that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonFloat64() *JsonFloat64 {
	return &JsonFloat64{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonFloat64 to Json
func (i *JsonFloat64) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%f", i.Value)), nil
}

// UnmarshalJSON converts from Json to JsonFloat64
func (i *JsonFloat64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp float64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of Float64 '0.0'
func (i *JsonFloat64) ValueOrDefault() float64 {
	return i.ValueOrDefaultValue(0)
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonFloat64) ValueOrDefaultValue(val float64) float64 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonFloat64 is valid and set
func (i *JsonFloat64) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonFloat64) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Float64ToString(i.Value)
}
