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

//goland:noinspection GoUnusedExportedFunction
func NewJsonFloat32(value float32) *JsonFloat32 {
	return &JsonFloat32{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NullJsonFloat32() *JsonFloat32 {
	return &JsonFloat32{
		Valid: false,
		Set:   true,
	}
}

func (i *JsonFloat32) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%f", i.Value)), nil
}

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

func (i *JsonFloat32) ValueOrDefault() float32 {
	return i.ValueOrDefaultValue(0.0)
}

func (i *JsonFloat32) ValueOrDefaultValue(val float32) float32 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

func (i *JsonFloat32) ValidAndSet() bool {
	return i.Set && i.Valid
}

func (i *JsonFloat32) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Float32ToString(i.Value)
}
