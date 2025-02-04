package main

import "testing"

func TestConvertToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "42"},
		{3.14, "3.14"},
		{"Golang", "Golang"},
		{true, "true"},
		{complex64(1 + 2i), "(1+2i)"},
	}

	for _, test := range tests {
		result := ConvertToString(test.input)
		if result != test.expected {
			t.Errorf("Ожидалось: %v, получено: %v", test.expected, result)
		}
	}
}
func TestSliceInStr(t *testing.T) {
	input := []string{"a", "b", "c"}
	expected := "abc"
	result := SliceInStr(input)
	if result != expected {
		t.Errorf("Ожидалось: %v, получено: %v", expected, result)
	}
}

func TestStrInSliceRune(t *testing.T) {
	input := "Helloworld"
	salt := "salt"
	expected := []rune("Hellosaltworld")
	result := StrInSliceRune(input, salt)
	if string(result) != string(expected) {
		t.Errorf("Ожидалось: %v, получено: %v", string(expected), string(result))
	}
}

func TestHashRunes(t *testing.T) {
	input := []rune("HelloWorld")
	expected := "872e4e50ce9990d8b041330c47c9ddd11bec6b503ae9386a99da8584e9bb12c4"
	result := HashRunes(input)
	if result != expected {
		t.Errorf("Ожидалось: %v, получено: %v", expected, result)
	}
}
