package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/umbrella-sh/um-common/ext"
)

type JsonFloat32 struct {
	Value float32 `json:"value"`
	Valid bool    `json:"valid"`
	Set   bool    `json:"set"`
}

// NewJsonFloat32 creates a new instance of JsonFloat32 with the 'value' defined
//
//goland:noinspection GoUnusedExportedFunction
func NewJsonFloat32(value float32) *JsonFloat32 {
	return &JsonFloat32{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

// NullJsonFloat32 creates a new instance of JsonFloat32 that will serialize to null
//
//goland:noinspection GoUnusedExportedFunction
func NullJsonFloat32() *JsonFloat32 {
	return &JsonFloat32{
		Valid: false,
		Set:   true,
	}
}

// MarshalJSON converts from JsonFloat32 to Json
func (i *JsonFloat32) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%f", i.Value)), nil
}

// UnmarshalJSON converts from Json to JsonFloat32
func (i *JsonFloat32) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp float32
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

// ValueOrDefault returns the value or the default value of Float32 '0.0'
func (i *JsonFloat32) ValueOrDefault() float32 {
	return i.ValueOrDefaultValue(0.0)
}

// ValueOrDefaultValue returns the value or the defined default value
func (i *JsonFloat32) ValueOrDefaultValue(val float32) float32 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

// ValidAndSet returns true if the JsonFloat32 is valid and set
func (i *JsonFloat32) ValidAndSet() bool {
	return i.Set && i.Valid
}

// ToString returns the value as a string
func (i *JsonFloat32) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Float32ToString(i.Value)
}
