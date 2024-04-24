package convert

func DeRef[T any](v *T) T {
	if v == nil {
		var out T

		return out
	}

	return *v
}

func Ptr[T any](v T) *T {
	return &v
}
