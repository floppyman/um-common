package jsons

import (
	"encoding/json"
	"fmt"

	"github.com/umbrella-sh/um-common/ext"
)

type JsonInt32 struct {
	Value int32 `json:"value"`
	Valid bool  `json:"valid"`
	Set   bool  `json:"set"`
}

//goland:noinspection GoUnusedExportedFunction
func NewJsonInt32(value int32) *JsonInt32 {
	return &JsonInt32{
		Value: value,
		Valid: true,
		Set:   true,
	}
}

//goland:noinspection GoUnusedExportedFunction
func NullJsonInt32() *JsonInt32 {
	return &JsonInt32{
		Valid: false,
		Set:   true,
	}
}

func (i *JsonInt32) MarshalJSON() ([]byte, error) {
	if !i.Set || (i.Set && !i.Valid) {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("%d", i.Value)), nil
}

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

func (i *JsonInt32) ValueOrDefault() int32 {
	return i.ValueOrDefaultValue(0)
}

func (i *JsonInt32) ValueOrDefaultValue(val int32) int32 {
	if i.Set && i.Valid {
		return i.Value
	}
	return val
}

func (i *JsonInt32) ValidAndSet() bool {
	return i.Set && i.Valid
}

func (i *JsonInt32) ToString() string {
	if !i.Set || (i.Set && !i.Valid) {
		return "null"
	}
	return ext.Int32ToString(i.Value)
}
