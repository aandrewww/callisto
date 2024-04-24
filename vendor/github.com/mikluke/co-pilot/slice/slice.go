package slice

import "github.com/mikluke/co-pilot/convert"

func Split[T any](s []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(s) {
		s, chunks = s[chunkSize:], append(chunks, s[0:chunkSize:chunkSize])
	}
	return append(chunks, s)
}

func Filter[T any](s []T, f func(v T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}

func Contains[T comparable](s []T, v T) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}

	return false
}

func Unique[T comparable](s []T) []T {
	m := make(map[T]struct{}, len(s))

	for _, v := range s {
		m[v] = struct{}{}
	}

	return convert.Map2Slice(m, func(k T, _ struct{}) T {
		return k
	})
}
