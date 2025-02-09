package jsons

import (
	"encoding/json"
	"fmt"
)

type JsonString struct {
	Value string `json:"value"`
	Valid bool   `json:"valid"`
	Set   bool   `json:"set"`
}

// NewJsonString creates a new instance of JsonString with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonString(value string) *JsonString {
	return &JsonString{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

// NullJsonString creates a new instance of JsonString that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonString() *JsonString {
	return &JsonString{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonString to Json
func (i *JsonString) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", i.Value)), nil
}

// UnmarshalJSON converts from Json to JsonString
func (i *JsonString) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of String '""'
func (i *JsonString) ValueOrDefault() string {
	return i.ValueOrDefaultValue("")
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonString) ValueOrDefaultValue(val string) string {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonString is valid and set
func (i *JsonString) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonString) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return i.Value
}
