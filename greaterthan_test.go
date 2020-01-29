package hamcrest_test

import (
	. "github.com/pepinns/go-hamcrest"
	"math"
	"testing"
)

// Uint > integer types
func TestUintTypesCanBeGreaterThanInt(t *testing.T) {
	gtVal := int(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(200), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanInt8(t *testing.T) {
	gtVal := int8(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(200), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanInt16(t *testing.T) {
	gtVal := int16(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(200), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanInt32(t *testing.T) {
	gtVal := int32(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(200), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanInt64(t *testing.T) {
	gtVal := int64(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(200), GreaterThan(gtVal))
}

// Int types > uint types
func TestIntTypesCanBeGreaterThanUInt(t *testing.T) {
	gtVal := uint(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanUInt8(t *testing.T) {
	gtVal := uint8(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanUInt16(t *testing.T) {
	gtVal := uint16(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanUInt32(t *testing.T) {
	gtVal := uint32(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanUInt64(t *testing.T) {
	gtVal := uint64(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}

// Int types > Int types
func TestIntTypesCanBeGreaterThanInt(t *testing.T) {
	gtVal := int(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanInt8(t *testing.T) {
	gtVal := int8(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanInt16(t *testing.T) {
	gtVal := int16(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanInt32(t *testing.T) {
	gtVal := int32(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}
func TestIntTypesCanBeGreaterThanInt64(t *testing.T) {
	gtVal := int64(1)
	Assert(t).That(int64(200), GreaterThan(gtVal))
	Assert(t).That(int32(200), GreaterThan(gtVal))
	Assert(t).That(int16(200), GreaterThan(gtVal))
	Assert(t).That(int8(100), GreaterThan(gtVal))
}

// Uint types > Uint types
func TestUintTypesCanBeGreaterThanUint(t *testing.T) {
	gtVal := uint(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(100), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanUint8(t *testing.T) {
	gtVal := uint8(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(100), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanUint16(t *testing.T) {
	gtVal := uint16(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(100), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanUint32(t *testing.T) {
	gtVal := uint32(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(100), GreaterThan(gtVal))
}
func TestUintTypesCanBeGreaterThanUint64(t *testing.T) {
	gtVal := uint64(1)
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(100), GreaterThan(gtVal))
}

func TestOverflowedInt8IsAlwaysGreaterThanInt8(t *testing.T) {
	gtVal := math.MaxInt8
	Assert(t).That(uint64(200), GreaterThan(gtVal))
	Assert(t).That(uint32(200), GreaterThan(gtVal))
	Assert(t).That(uint16(200), GreaterThan(gtVal))
	Assert(t).That(uint8(200), GreaterThan(gtVal))
}
func TestMaxInt64IsNotGreaterThanMaxUint64(t *testing.T) {
	Assert(t).That(uint64(math.MaxUint64), GreaterThan(math.MaxInt64))
	Assert(t).That(int64(math.MaxInt64), Not(GreaterThan(uint64(math.MaxUint64))))
}
