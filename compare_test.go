package jgptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareString(t *testing.T) {
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
		assert.Equal(t, c.expect, CompareString(c.a, c.b), "case %v", i)
	}
}

func TestCompareUint64(t *testing.T) {
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
		assert.Equal(t, c.expect, CompareUint64(c.a, c.b), "case %v", i)
	}
}
