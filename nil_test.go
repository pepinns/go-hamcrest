package hamcrest_test

import (
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

func TestCanMatchAgainstNil(t *testing.T) {
	Assert(t).That(nil, IsNil())
}
func TestCanMatchAgainstNilPtrInterface(t *testing.T) {
	type someStruct struct {
	}
	fun := func() *someStruct {
		return nil
	}

	Assert(t).That(fun(), IsNil())
}
