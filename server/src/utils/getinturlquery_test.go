package utils

import (
	"testing"
)

func TestGetIntURLQuery(t *testing.T) {
	testKeyIsNotInMap(t)
	testEmptyValue(t)
	testInvalidIntStrings(t)
	testValidIntStrings(t)
}

func assertErr(t *testing.T, err error, want string) {
	if err == nil {
		t.Fail()
	} else {
		got := err.Error()
		if got != want {
			t.Errorf("got %#v; want %#v", got, want)
		}
	}
}

func testEmptyValue(t *testing.T) {
	m := map[string][]string{
		"key": []string{},
	}
	_, err := GetIntURLQuery(m, "key")
	assertErr(t, err, "getIntURLQuery: empty value")
}

func testInvalidIntStrings(t *testing.T) {
	m := map[string][]string{
		"a": []string{""},
		"b": []string{"invalid"},
	}
	wants := map[string]string{
		"a": "strconv.Atoi: parsing \"\": invalid syntax",
		"b": "strconv.Atoi: parsing \"invalid\": invalid syntax",
	}
	for key := range m {
		_, err := GetIntURLQuery(m, key)
		assertErr(t, err, wants[key])
	}
}

func testKeyIsNotInMap(t *testing.T) {
	m := map[string][]string{}
	_, err := GetIntURLQuery(m, "key")
	assertErr(t, err, "getIntURLQuery: key isn't in map")
}

func testValidIntStrings(t *testing.T) {
	m := map[string][]string{
		"a": []string{"-1"},
		"b": []string{"0"},
		"c": []string{"1"},
		"d": []string{"286"},
	}
	wants := map[string]int{
		"a": -1,
		"b": 0,
		"c": 1,
		"d": 286,
	}
	for key := range m {
		got, err := GetIntURLQuery(m, key)
		if err != nil {
			t.Fail()
		}
		if got != wants[key] {
			t.Errorf("got %#v; want: %#v", got, wants[key])
		}
	}
}
