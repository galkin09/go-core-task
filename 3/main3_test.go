package main

import (
	"testing"
)

func TestStringIntMap(t *testing.T) {
	m := NewStringIntMap(make(map[string]int))

	m.Add("1", 1)
	if !m.Exists("1") {
		t.Error("ключа 1 не существует")
	}

	v, ok := m.Get("1")
	if !ok || v != 1 {
		t.Errorf("для ключа 1 значение не равно 1")
	}

	m.Remove("1")
	if m.Exists("1") {
		t.Errorf("ключ 1 не удален")
	}

	m.Add("2", 2)
	m.Add("3", 3)
	copyMap := m.Copy()

	if len(copyMap) != 2 {
		t.Errorf("размер скопированный мапы не 2")
	}
	if copyMap["2"] != 2 || copyMap["3"] != 3 {
		t.Errorf("скопированы неверные элементы")
	}
}
