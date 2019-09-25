package hamcrest_test

import (
	"bytes"
	. "hamcrest"
	"io"
	"testing"
)

func TestStringContainsMatchesStringInsideString(t *testing.T) {
	Assert(t).That("some long string", Contains("long"))
}
func TestStringNotContainsMatchesStringInsideString(t *testing.T) {
	Assert(t).That("some long string", Not(Contains("extra longlong")))
}

type TestReader struct {
	inner io.Reader
}

func (me *TestReader) Read(b []byte) (int, error) {
	return me.inner.Read(b)
}

func TestIoReaderContainsMatchString(t *testing.T) {
	b := bytes.NewBufferString("Some data to read as a buffer")
	Assert(t).That(io.Reader(&TestReader{b}), Contains("read as"))
}
func TestIoReaderErrorMessageIsClear(t *testing.T) {
	b := bytes.NewBufferString("Some data to read as a buffer")
	AssertFailureMessage(t, io.Reader(&TestReader{b}), Contains("read as"), Equals(`"Some data to read as a buffer" contains string "read as"`))
}
func TestIoReaderErrorMessageIsClearForNegative(t *testing.T) {
	b := bytes.NewBufferString("Some data to read as a buffer")
	AssertFailureMessage(t, io.Reader(&TestReader{b}), Contains("NOT MATCH read as"), Equals(`"Some data to read as a buffer" does not contain string "NOT MATCH read as"`))
}

func TestContainsCanTakeMultipleArgsAndTurnIntoAllOf(t *testing.T) {
	Assert(t).That([]string{"somestring", "otherstring"}, Contains("otherstring", "somestring"))
}

func TestStringSliceCanContainString(t *testing.T) {
	Assert(t).That([]string{"somestring", "otherstring"}, Contains("otherstring"))
}
func TestStringSliceMessageIsClear(t *testing.T) {
	AssertFailureMessage(t, []string{"item1", "item2"}, Contains("item2"), Equals(`matched items [
  [1] "item2" is equal to "item2"
]`))
	AssertFailureMessage(t, []string{"item1", "item2"}, Contains("item3"), Equals(`failed to match [
  [0] "item1" is not equal to "item3"
  [1] "item2" is not equal to "item3"
]`))
}
func TestStringSliceCanNotContainString(t *testing.T) {
	Assert(t).That([]string{"somestring", "otherstring"}, Not(Contains("somestring otherstring")))
}

func TestStringSliceCanNotContainInt(t *testing.T) {
	Assert(t).That([]string{"somestring", "otherstring"}, Not(Contains(38)))
}

func TestIntegerSliceCanContainInt(t *testing.T) {
	Assert(t).That([]int{34, 38}, Contains(38))
}

func TestInt32SliceCanContainInt(t *testing.T) {
	Assert(t).That([]int32{int32(34), int32(38)}, Contains(38))
}
func TestInt32SliceCanContainInt16(t *testing.T) {
	Assert(t).That([]int32{int32(34), int32(38)}, Contains(int16(38)))
}
func TestIntegerSliceCanNotContainInt(t *testing.T) {
	Assert(t).That([]int{34, 38}, Not(Contains(438)))
}
