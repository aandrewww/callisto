package convert

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromProtoTimestamp(v *timestamppb.Timestamp) *time.Time {
	if v == nil {
		return nil
	}

	return Ptr(v.AsTime())
}

func ToProtoTimestamp(v *time.Time) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}

	return timestamppb.New(*v)
}
