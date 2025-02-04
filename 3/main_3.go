package main

import "fmt"

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap(m map[string]int) *StringIntMap {
	return &StringIntMap{m: m}
}

func (sm *StringIntMap) Add(k string, v int) {
	sm.m[k] = v
}

func (sm *StringIntMap) Remove(k string) {
	delete(sm.m, k)
}

func (sm *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range sm.m {
		newMap[k] = v
	}
	return newMap
}

func (sm *StringIntMap) Exists(k string) bool {
	_, ok := sm.m[k]
	return ok
}

func (sm *StringIntMap) Get(k string) (int, bool) {
	v, ok := sm.m[k]
	return v, ok
}

func main() {
	myMap := NewStringIntMap(make(map[string]int))

	myMap.Add("1", 1)
	myMap.Add("2", 2)
	myMap.Add("3", 3)

	fmt.Println(myMap)

	myMap.Remove("2")

	copyMap := myMap.Copy()
	fmt.Println(copyMap)

	fmt.Println(myMap.Exists("1"))
	fmt.Println(myMap.Exists("2"))
}
