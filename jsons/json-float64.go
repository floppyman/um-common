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

//goland:noinspection GoUnusedExportedFunction
func NewJsonFloat64(value float64) *JsonFloat64 {
	return &JsonFloat64{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NullJsonFloat64() *JsonFloat64 {
	return &JsonFloat64{
		Valid: false,
		Set:   true,
	}
}

func (i *JsonFloat64) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%f", i.Value)), nil
}

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

func (i *JsonFloat64) ValueOrDefault() float64 {
	return i.ValueOrDefaultValue(0)
}

func (i *JsonFloat64) ValueOrDefaultValue(val float64) float64 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

func (i *JsonFloat64) ValidAndSet() bool {
	return i.Set && i.Valid
}

func (i *JsonFloat64) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Float64ToString(i.Value)
}
