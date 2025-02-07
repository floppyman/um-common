package jsons

import (
	"encoding/json"
	"fmt"
)

type JsonObject[TValue any] struct {
	Value TValue `json:"value"`
	Valid bool   `json:"valid"`
	Set   bool   `json:"set"`
}

//goland:noinspection GoUnusedExportedFunction
func NewJsonObject[TValue any](value TValue) *JsonObject[TValue] {
	return &JsonObject[TValue]{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NullJsonObject[TValue any]() *JsonObject[TValue] {
	return &JsonObject[TValue]{
		Valid: false,
		Set:   true,
	}
}

func (i *JsonObject[TValue]) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	bytes, err := json.Marshal(i.Value)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (i *JsonObject[TValue]) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp TValue
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *JsonObject[TValue]) ValueOrDefaultValue(val TValue) TValue {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

func (i *JsonObject[TValue]) ValidAndSet() bool {
	return i.Set && i.Valid
}

func (i *JsonObject[TValue]) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return fmt.Sprintf("\"%v\"", i.Value)
}
