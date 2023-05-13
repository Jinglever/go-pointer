package jgptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareStringPointer(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *string
		expect bool
	}{
		{nil, nil, true},
		{nil, NewString("a"), false},
		{NewString("a"), nil, false},
		{NewString("a"), NewString("a"), true},
		{NewString("a"), NewString("b"), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareStringPointer(c.a, c.b), "case %v", i)
	}
}

func TestCompareUint64Pointer(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *uint64
		expect bool
	}{
		{nil, nil, true},
		{nil, NewUint64(1), false},
		{NewUint64(1), nil, false},
		{NewUint64(1), NewUint64(1), true},
		{NewUint64(1), NewUint64(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareUint64Pointer(c.a, c.b), "case %v", i)
	}
}
