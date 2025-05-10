package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/floppyman/um-common/ext"
)

type JsonInt64 struct {
	Value int64 `json:"value"`
	Valid bool  `json:"valid"`
	Set   bool  `json:"set"`
}

// NewJsonInt64 creates a new instance of JsonInt64 with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonInt64(value int64) *JsonInt64 {
	return &JsonInt64{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

// NullJsonInt64 creates a new instance of JsonInt64 that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonInt64() *JsonInt64 {
	return &JsonInt64{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonInt64 to Json
func (i *JsonInt64) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%d", i.Value)), nil
}

// UnmarshalJSON converts from Json to JsonInt64
func (i *JsonInt64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of Int64 '0'
func (i *JsonInt64) ValueOrDefault() int64 {
	return i.ValueOrDefaultValue(0)
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonInt64) ValueOrDefaultValue(val int64) int64 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonInt64 is valid and set
func (i *JsonInt64) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonInt64) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Int64ToString(i.Value)
}
