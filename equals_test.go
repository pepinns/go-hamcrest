package hamcrest_test

import (
	. "github.com/pepinns/go-hamcrest"
	"testing"
)

func TestStringEqualsReturnsTrueOnEquality(t *testing.T) {
	eq := &StringEqualsMatcher{MatchValue: "other"}

	result := eq.Match("other")

	if !result.Matched() {
		t.Fatal("expected 2 strings to match")
	}
}
func TestStringEqualsReturnsFalseOnEquality(t *testing.T) {
	eq := &StringEqualsMatcher{MatchValue: "other"}

	result := eq.Match("something different")

	if result.Matched() {
		t.Fatal("expected 2 strings to NOT match")
	}
}

func TestStringNotEqualsMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, "stringToMatch", Equals("otherString"), Equals(`"stringToMatch" is not equal to "otherString"`))
}
func TestStringEqualsMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, "stringToMatch", Equals("stringToMatch"), Equals(`"stringToMatch" is equal to "stringToMatch"`))
}

func TestStringEqualsWithAsserter(t *testing.T) {
	Assert(t).That("prefix", Equals("prefix"))
}

func TestIntegerEqualsAnotherInteger(t *testing.T) {
	Assert(t).That(23, Equals(23))
}
func TestIntEqualsMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, 42, Equals(42), Equals("'42' is equal to '42'"))
}
func TestIntNotEqualsMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, 45, Equals(42), Equals("'45' is not equal to '42'"))
}
func TestInt8EqualsAnotherInteger(t *testing.T) {
	Assert(t).That(int8(23), Equals(23))
}
func TestInt16EqualsAnotherInteger(t *testing.T) {
	Assert(t).That(int16(23), Equals(23))
}
func TestInt32EqualsAnotherInteger(t *testing.T) {
	Assert(t).That(int32(23), Equals(23))
}
func TestInt32EqualsAnotherInt32(t *testing.T) {
	Assert(t).That(int32(23), Equals(int32(23)))
}
func TestUInt32EqualsMatcherInt(t *testing.T) {
	Assert(t).That(uint32(23), Equals(23))
}
func TestUInt64EqualsMatcherInt(t *testing.T) {
	Assert(t).That(uint64(23), Equals(23))
}

func TestUInt64EqualsMatcherUInt(t *testing.T) {
	Assert(t).That(uint64(23), Equals(uint64(23)))
}
func TestUInt64EqualsMatcherUInt32(t *testing.T) {
	Assert(t).That(uint64(23), Equals(uint32(23)))
}
func TestUintInMatcherWontAllowOverflows(t *testing.T) {
	Assert(t).That(-45, Not(Equals(uint64(18446744073709551571))))
	AssertFailureMessage(t, -45, Equals(uint64(18446744073709551571)), Equals("'-45' is < 0 and cannot be converted to uint"))
}
func TestUIntEqualsMessageIsClear(t *testing.T) {
	Assert(t).That(42, Equals(uint64(42)))

	AssertFailureMessage(t, 42, Equals(uint64(42)), Equals("'42' is equal to '42'"))
}
func TestUIntNotEqualsMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, 45, Equals(uint64(42)), Equals("'45' is not equal to '42'"))
}

func TestFloatEqualsAnotherFloat(t *testing.T) {
	Assert(t).That(23.41, Equals(23.41))
}
func TestFloat32EqualsAnotherFloat(t *testing.T) {
	Assert(t).That(float32(23.41), Equals(23.41))
}
func TestFloatMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, 23.41, Equals(55.12), Equals(`"23.4100" is not equal to "55.1200"`))
}

func TestBoolEqualsAnotherBool(t *testing.T) {
	Assert(t).That(false, Equals(false))
}
func TestByteslicecanequalbyteslice(t *testing.T) {
	Assert(t).That([]byte("somestuff"), Equals([]byte("somestuff")))
}
func TestByteSliceEqualsMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, []byte("deadbeef"), Equals([]byte("deadbeef")), Equals(`"6465616462656566" is equal to "6465616462656566"`))
	AssertFailureMessage(t, []byte("deadbeefff"), Equals([]byte("deadbeef")), Equals(`"64656164626565666666" is not equal to "6465616462656566"`))
}

func TestStringNotEqualsAnotherString(t *testing.T) {
	Assert(t).That("somestring", Not(Equals("xxxsomestring")))
}

func TestEqualsCanCompareMaps(t *testing.T) {
	sourceMap := make(map[string]string)
	testMap := make(map[string]string)

	sourceMap["foo"] = "bar"
	testMap["foo"] = "bar"
	Assert(t).That(sourceMap, Equals(testMap))
}
