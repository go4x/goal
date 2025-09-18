package pointer

func To[T any](v T) *T {
	return &v
}

func From[T any](v *T) T {
	return *v
}

func ToSlice[T any](v []T) []*T {
	var ret []*T
	for _, v := range v {
		ret = append(ret, To(v))
	}
	return ret
}

func FromSlice[T any](v []*T) []T {
	var ret []T
	for _, v := range v {
		ret = append(ret, From(v))
	}
	return ret
}
