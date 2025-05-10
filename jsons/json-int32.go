package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/floppyman/um-common/ext"
)

type JsonInt32 struct {
	Value int32 `json:"value"`
	Valid bool  `json:"valid"`
	Set   bool  `json:"set"`
}

// NewJsonInt32 creates a new instance of JsonInt32 with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonInt32(value int32) *JsonInt32 {
	return &JsonInt32{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

// NullJsonInt32 creates a new instance of JsonInt32 that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonInt32() *JsonInt32 {
	return &JsonInt32{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonInt32 to Json
func (i *JsonInt32) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%d", i.Value)), nil
}

// UnmarshalJSON converts from Json to JsonInt32
func (i *JsonInt32) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int32
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of Int32 '0'
func (i *JsonInt32) ValueOrDefault() int32 {
	return i.ValueOrDefaultValue(0)
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonInt32) ValueOrDefaultValue(val int32) int32 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonInt32 is valid and set
func (i *JsonInt32) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonInt32) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Int32ToString(i.Value)
}
