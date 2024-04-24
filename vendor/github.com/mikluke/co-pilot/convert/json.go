package convert

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func FromJSON[T any](b []byte) (*T, error) {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &v, nil
}

func CompactJSON(b []byte) ([]byte, error) {
	if b == nil {
		return nil, nil
	}

	buf := bytes.NewBuffer(make([]byte, 0, len(b)))

	if err := json.Compact(buf, b); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
