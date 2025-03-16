package main

import "fmt"

// StringIntMap — структура для хранения пар "строка — целое число"
type StringIntMap struct {
	data map[string]int
}

func (s *StringIntMap) String() {
	for k, v := range s.data {
        fmt.Printf("%s: %d\n", k, v)
    }
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

/*
func (s *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(s.data))

	for k, v := range s.data {
		newMap[k] = v
	}

	return newMap
}
*/

func (s *StringIntMap) Copy() StringIntMap {
	newMap := StringIntMap{
		data: make(map[string]int, len(s.data)),
	}

    for k, v := range s.data {
        newMap.data[k] = v
    }

    return newMap
}


func (s *StringIntMap) Exists(key string) bool {
	_, exists := s.data[key]
    return exists
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, exists := s.data[key]
    return value, exists
}
