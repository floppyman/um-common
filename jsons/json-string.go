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

//goland:noinspection GoUnusedExportedFunction
func NewJsonString(value string) *JsonString {
	return &JsonString{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NullJsonString() *JsonString {
	return &JsonString{
		Valid: false,
		Set:   true,
	}
}

func (i *JsonString) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", i.Value)), nil
}

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

func (i *JsonString) ValueOrDefault() string {
	return i.ValueOrDefaultValue("")
}

func (i *JsonString) ValueOrDefaultValue(val string) string {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

func (i *JsonString) ValidAndSet() bool {
	return i.Set && i.Valid
}

func (i *JsonString) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return i.Value
}
