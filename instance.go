package jgptr

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