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

func TestCompareInt64(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *int64
		expect bool
	}{
		{nil, nil, true},
		{nil, NewInt64(1), false},
		{NewInt64(1), nil, false},
		{NewInt64(1), NewInt64(1), true},
		{NewInt64(1), NewInt64(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareInt64(c.a, c.b), "case %v", i)
	}
}

func TestCompareUint8(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *uint8
		expect bool
	}{
		{nil, nil, true},
		{nil, NewUint8(1), false},
		{NewUint8(1), nil, false},
		{NewUint8(1), NewUint8(1), true},
		{NewUint8(1), NewUint8(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareUint8(c.a, c.b), "case %v", i)
	}
}

func TestCompareInt8(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *int8
		expect bool
	}{
		{nil, nil, true},
		{nil, NewInt8(1), false},
		{NewInt8(1), nil, false},
		{NewInt8(1), NewInt8(1), true},
		{NewInt8(1), NewInt8(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareInt8(c.a, c.b), "case %v", i)
	}
}

func TestCompareUint16(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *uint16
		expect bool
	}{
		{nil, nil, true},
		{nil, NewUint16(1), false},
		{NewUint16(1), nil, false},
		{NewUint16(1), NewUint16(1), true},
		{NewUint16(1), NewUint16(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareUint16(c.a, c.b), "case %v", i)
	}
}

func TestCompareInt16(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *int16
		expect bool
	}{
		{nil, nil, true},
		{nil, NewInt16(1), false},
		{NewInt16(1), nil, false},
		{NewInt16(1), NewInt16(1), true},
		{NewInt16(1), NewInt16(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareInt16(c.a, c.b), "case %v", i)
	}
}

func TestCompareUint32(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *uint32
		expect bool
	}{
		{nil, nil, true},
		{nil, NewUint32(1), false},
		{NewUint32(1), nil, false},
		{NewUint32(1), NewUint32(1), true},
		{NewUint32(1), NewUint32(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareUint32(c.a, c.b), "case %v", i)
	}
}

func TestCompareInt32(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *int32
		expect bool
	}{
		{nil, nil, true},
		{nil, NewInt32(1), false},
		{NewInt32(1), nil, false},
		{NewInt32(1), NewInt32(1), true},
		{NewInt32(1), NewInt32(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareInt32(c.a, c.b), "case %v", i)
	}
}

func TestCompareFloat64(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *float64
		expect bool
	}{
		{nil, nil, true},
		{nil, NewFloat64(1), false},
		{NewFloat64(1), nil, false},
		{NewFloat64(1), NewFloat64(1), true},
		{NewFloat64(1), NewFloat64(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareFloat64(c.a, c.b), "case %v", i)
	}
}

func TestCompareFloat32(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *float32
		expect bool
	}{
		{nil, nil, true},
		{nil, NewFloat32(1), false},
		{NewFloat32(1), nil, false},
		{NewFloat32(1), NewFloat32(1), true},
		{NewFloat32(1), NewFloat32(2), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareFloat32(c.a, c.b), "case %v", i)
	}
}

func TestCompareBool(t *testing.T) {
	// use case and assert
	cases := []struct {
		a, b   *bool
		expect bool
	}{
		{nil, nil, true},
		{nil, NewBool(true), false},
		{NewBool(true), nil, false},
		{NewBool(true), NewBool(true), true},
		{NewBool(true), NewBool(false), false},
	}
	for i, c := range cases {
		assert.Equal(t, c.expect, CompareBool(c.a, c.b), "case %v", i)
	}
}
