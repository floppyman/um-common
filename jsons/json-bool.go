package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/umbrella-sh/um-common/ext"
)

type JsonBool struct {
	Value bool `json:"value"`
	Valid bool `json:"valid"`
	Set   bool `json:"set"`
}

//goland:noinspection GoUnusedExportedFunction
func NewJsonBool(value bool) *JsonBool {
	return &JsonBool{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NullJsonBool() *JsonBool {
	return &JsonBool{
		Valid: false,
		Set:   true,
	}
}

func (i *JsonBool) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%t", i.Value)), nil
}

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

func (i *JsonBool) ValueOrDefault() bool {
	return i.ValueOrDefaultValue(false)
}

func (i *JsonBool) ValueOrDefaultValue(val bool) bool {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

func (i *JsonBool) ValidAndSet() bool {
	return i.Set && i.Valid
}

func (i *JsonBool) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Iif(i.Value, "true", "false")
}
