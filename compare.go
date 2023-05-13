package jgptr

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
