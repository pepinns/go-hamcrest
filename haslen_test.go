package hamcrest_test

import (
	. "hamcrest"
	"testing"
)

func TestHasLenCanMatchStringLength(t *testing.T) {
	Assert(t).That("one", HasLen(3))
}

func TestHasLenCanMatchSliceLength(t *testing.T) {
	tt := make([]int, 10)
	Assert(t).That(tt, AllOf(
		HasLen(10),
		HasLen(Equals(10)),
		HasLen(Not(Equals(34))),
	))
}

func TestHasLenResponseIsClear(t *testing.T) {
	// AssertFailureMessage(t, "one", HasLen(5), Equals("Fred"))
	// AssertFailureMessage(t, "one", Not(HasLen(5)), Equals("Fred"))
}
