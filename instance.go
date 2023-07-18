package jgptr

import "time"

// GetString get string by pointer
// return "" if pointer is nil
func GetString(pointer *string) string {
	if pointer == nil {
		return ""
	}
	return *pointer
}

// GetUint64 get uint64 by pointer
// return 0 if pointer is nil
func GetUint64(pointer *uint64) uint64 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetInt64 get int64 by pointer
// return 0 if pointer is nil
func GetInt64(pointer *int64) int64 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetUint8 get uint8 by pointer
// return 0 if pointer is nil
func GetUint8(pointer *uint8) uint8 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetInt8 get int8 by pointer
// return 0 if pointer is nil
func GetInt8(pointer *int8) int8 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetUint16 get uint16 by pointer
// return 0 if pointer is nil
func GetUint16(pointer *uint16) uint16 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetInt16 get int16 by pointer
// return 0 if pointer is nil
func GetInt16(pointer *int16) int16 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetUint32 get uint32 by pointer
// return 0 if pointer is nil
func GetUint32(pointer *uint32) uint32 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetInt32 get int32 by pointer
// return 0 if pointer is nil
func GetInt32(pointer *int32) int32 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetFloat64 get float64 by pointer
// return 0 if pointer is nil
func GetFloat64(pointer *float64) float64 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetFloat32 get float32 by pointer
// return 0 if pointer is nil
func GetFloat32(pointer *float32) float32 {
	if pointer == nil {
		return 0
	}
	return *pointer
}

// GetBool get bool by pointer
// return false if pointer is nil
func GetBool(pointer *bool) bool {
	if pointer == nil {
		return false
	}
	return *pointer
}

// GetTime get time.Time by pointer
// return zero time if pointer is nil
func GetTime(pointer *time.Time) time.Time {
	if pointer == nil {
		return time.Time{}
	}
	return *pointer
}
