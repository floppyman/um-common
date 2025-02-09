package web

import "testing"

func Test_QueryStringBuilder(t *testing.T) {
	var items = make(map[string]string)
	items["key1"] = "value1"
	items["key2"] = "value2"
	items["key3"] = "value3"

	expectedRes := "key1=value1&key2=value2&key3=value3"
	res := QueryStringBuilder(items)
	if res != expectedRes {
		t.Errorf("Expected '%s' to be '%s'", res, expectedRes)
	}
}
