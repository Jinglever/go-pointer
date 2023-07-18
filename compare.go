package jgptr

import "time"

// CompareString compare string pointer
// return true if a and b are both nil or a == b
func CompareString(a, b *string) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareUint64 compare uint64 pointer
// return true if a and b are both nil or a == b
func CompareUint64(a, b *uint64) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareInt64 compare int64 pointer
// return true if a and b are both nil or a == b
func CompareInt64(a, b *int64) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareUint8 compare uint8 pointer
// return true if a and b are both nil or a == b
func CompareUint8(a, b *uint8) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareInt8 compare int8 pointer
// return true if a and b are both nil or a == b
func CompareInt8(a, b *int8) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareUint16 compare uint16 pointer
// return true if a and b are both nil or a == b
func CompareUint16(a, b *uint16) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareInt16 compare int16 pointer
// return true if a and b are both nil or a == b
func CompareInt16(a, b *int16) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareUint32 compare uint32 pointer
// return true if a and b are both nil or a == b
func CompareUint32(a, b *uint32) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareInt32 compare int32 pointer
// return true if a and b are both nil or a == b
func CompareInt32(a, b *int32) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareFloat64 compare float64 pointer
// return true if a and b are both nil or a == b
func CompareFloat64(a, b *float64) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareFloat32 compare float32 pointer
// return true if a and b are both nil or a == b
func CompareFloat32(a, b *float32) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareBool compare bool pointer
// return true if a and b are both nil or a == b
func CompareBool(a, b *bool) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareTime compare time.Time pointer
// return true if a and b are both nil or a == b
func CompareTime(a, b *time.Time) bool {
	if a != nil && b != nil {
		return a.Equal(*b)
	}
	if a == nil && b == nil {
		return true
	}
	return false
}
