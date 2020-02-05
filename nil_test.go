package hamcrest_test

import (
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

type someStruct struct {
}

func TestCanMatchAgainstNil(t *testing.T) {
	Assert(t).That(nil, IsNil())
}
func TestCanMatchAgainstNilPtrInterface(t *testing.T) {
	fun := func() *someStruct {
		return nil
	}

	Assert(t).That(fun(), IsNil())
}

func TestDoesntMatchStructValue(t *testing.T) {
	val := &someStruct{}
	Assert(t).That(val, Not(IsNil()))
	val2 := someStruct{}
	Assert(t).That(val2, Not(IsNil()))
}
