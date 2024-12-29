package main

func flattenArray[T any](arr [][]T) []T {
	var flat []T
	for _, innerArray := range arr {
		flat = append(flat, innerArray...)
	}
	return flat
}

func getMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}