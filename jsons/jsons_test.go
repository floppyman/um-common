package jsons

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJsonBoolMarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonBool(true)
	var objUnmarshal JsonBool

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonBoolMarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonBool()
	var objUnmarshal JsonBool

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonFloat32MarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonFloat32(10.10)
	var objUnmarshal JsonFloat32

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonFloat32MarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonFloat32()
	var objUnmarshal JsonFloat32

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonFloat64MarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonFloat64(10.10)
	var objUnmarshal JsonFloat64

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonFloat64MarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonFloat64()
	var objUnmarshal JsonFloat64

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonInt32MarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonInt32(10)
	var objUnmarshal JsonInt32

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonInt32MarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonInt32()
	var objUnmarshal JsonInt32

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonInt64MarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonInt64(10)
	var objUnmarshal JsonInt64

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonInt64MarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonInt64()
	var objUnmarshal JsonInt64

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonObjectMarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonObject(JsonObjectTest{
		Number: 10,
		String: "string",
		Bool:   true,
	})
	var objUnmarshal JsonObject[JsonObjectTest]

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonObjectMarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonObject[JsonObjectTest]()
	var objUnmarshal JsonObject[JsonObjectTest]

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonStringMarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonString("string")
	var objUnmarshal JsonString

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonStringMarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonString()
	var objUnmarshal JsonString

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

func TestJsonTimeMarshalUnmarshal(t *testing.T) {
	objMarshal := NewJsonTime(time.Now())
	var objUnmarshal JsonTime

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value.String() == objUnmarshal.Value.String() {
		return
	}
	t.FailNow()
}

func TestJsonTimeMarshalUnmarshalNull(t *testing.T) {
	objMarshal := NullJsonTime()
	var objUnmarshal JsonTime

	bytes, err := json.Marshal(objMarshal)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &objUnmarshal)
	if err != nil {
		t.Fatal(err)
	}

	if objMarshal.Set == objUnmarshal.Set && objMarshal.Valid == objUnmarshal.Valid && objMarshal.Value == objUnmarshal.Value {
		return
	}
	t.FailNow()
}

type JsonObjectTest struct {
	Number int    `json:"number"`
	String string `json:"string"`
	Bool   bool   `json:"bool"`
}
