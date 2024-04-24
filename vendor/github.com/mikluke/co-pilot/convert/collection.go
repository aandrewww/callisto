package convert

func Slice[From any, To any](s []From, f func(v From) To) []To {
	if s == nil {
		return nil
	}

	out := make([]To, len(s))
	for i, v := range s {
		out[i] = f(v)
	}

	return out
}

func Map[FromKey comparable, FromValue any, ToKey comparable, ToValue any](
	m map[FromKey]FromValue,
	f func(k FromKey, v FromValue) (ToKey, ToValue)) map[ToKey]ToValue {
	if m == nil {
		return nil
	}

	out := make(map[ToKey]ToValue, len(m))
	for k, v := range m {
		ck, cv := f(k, v)
		out[ck] = cv
	}

	return out
}

func Slice2Map[From any, ToKey comparable, ToValue any](s []From, f func(v From) (ToKey, ToValue)) map[ToKey]ToValue {
	m := make(map[ToKey]ToValue, len(s))
	for i := range s {
		k, v := f(s[i])
		m[k] = v
	}

	return m
}

func Map2Slice[FromKey comparable, FromValue any, To any](m map[FromKey]FromValue, f func(k FromKey, v FromValue) To) []To {
	s := make([]To, 0, len(m))
	for k, v := range m {
		s = append(s, f(k, v))
	}

	return s
}
