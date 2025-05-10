package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/floppyman/um-common/ext"
)

// JsonBool defines a struct for handling json values that can be null
type JsonBool struct {
	Value bool `json:"value"`
	Valid bool `json:"valid"`
	Set   bool `json:"set"`
}

// NewJsonBool creates a new instance of JsonBool with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonBool(value bool) *JsonBool {
	return &JsonBool{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

// NullJsonBool creates a new instance of JsonBool that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonBool() *JsonBool {
	return &JsonBool{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonBool to Json
func (i *JsonBool) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%t", i.Value)), nil
}

// UnmarshalJSON converts from Json to JsonBool
func (i *JsonBool) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of bool 'false'
func (i *JsonBool) ValueOrDefault() bool {
	return i.ValueOrDefaultValue(false)
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonBool) ValueOrDefaultValue(val bool) bool {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonBool is valid and set
func (i *JsonBool) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonBool) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Iif(i.Value, "true", "false")
}
