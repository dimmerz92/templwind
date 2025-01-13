package lib

// Coalesce returns v1 if it is not the zero value for it's type, otherwise v2.
func Coalesce[T comparable](v1, v2 T) T {
	var zero T
	if v1 != zero {
		return v1
	}
	return v2
}

// IIF is an inline if statement. If the condition is true, it returns v1, otherwise v2.
func IIF[T any](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}
	return v2
}
