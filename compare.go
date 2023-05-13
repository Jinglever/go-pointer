package jgptr

// CompareStringPointer compare string pointer
// return true if a and b are both nil or a == b
func CompareStringPointer(a, b *string) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}

// CompareUint64Pointer compare uint64 pointer
// return true if a and b are both nil or a == b
func CompareUint64Pointer(a, b *uint64) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	if a == nil && b == nil {
		return true
	}
	return false
}
