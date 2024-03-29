package jgptr

import "time"

func NewString(s string) *string {
	return &s
}

func NewUint64(i uint64) *uint64 {
	return &i
}

func NewInt64(i int64) *int64 {
	return &i
}

func NewUint8(i uint8) *uint8 {
	return &i
}

func NewInt8(i int8) *int8 {
	return &i
}

func NewUint16(i uint16) *uint16 {
	return &i
}

func NewInt16(i int16) *int16 {
	return &i
}

func NewUint32(i uint32) *uint32 {
	return &i
}

func NewInt32(i int32) *int32 {
	return &i
}

func NewFloat64(f float64) *float64 {
	return &f
}

func NewFloat32(f float32) *float32 {
	return &f
}

func NewBool(b bool) *bool {
	return &b
}

func NewTime(t time.Time) *time.Time {
	return &t
}
