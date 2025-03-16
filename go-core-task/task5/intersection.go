package main

// FindIntersection проверяет наличие пересечений и возвращает пересекающиеся элементы
func FindIntersection(a, b []int) (bool, []int) {
	m := make(map[int]struct{}, len(a))

	for _, v := range a {
        m[v] = struct{}{}
    }

	intersection := make([]int, 0, len(a))
	found := false
	for _, v := range b {
        if _, ok := m[v]; ok {
            intersection = append(intersection, v)
			found = true
        }
    }

	return found, intersection
}