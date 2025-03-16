package main

// Difference возвращает слайс строк, содержащий элементы,
// которые есть в первом слайсе, но отсутствуют во втором
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]struct{})

	for _, v := range slice2 {
		m[v] = struct{}{}
	}

	result := make([]string, 0, len(slice1))

	for _, v := range slice1 {
        if _, exists := m[v]; !exists {
            result = append(result, v)
        }
    }

	return result
}