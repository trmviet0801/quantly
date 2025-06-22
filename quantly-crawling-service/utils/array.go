package utils

func ToPointerArray[T any](arr []T) []*T {
	var result []*T
	for _, item := range arr {
		result = append(result, &item)
	}
	return result
}
