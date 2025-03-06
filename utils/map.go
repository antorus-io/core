package utils

func Map[T, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))

	for i, item := range slice {
		result[i] = transform(item)
	}

	return result
}
