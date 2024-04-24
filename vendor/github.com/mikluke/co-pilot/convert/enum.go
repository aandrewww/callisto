package convert

import "fmt"

func Enum[FromType comparable, ToType any](v FromType, m map[FromType]ToType) (ToType, error) {
	out, ok := m[v]
	if !ok {
		return out, fmt.Errorf("mapping for %v not found", v)
	}

	return out, nil
}
