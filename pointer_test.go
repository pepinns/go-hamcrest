package hamcrest_test

import (
	. "hamcrest"
	"testing"
)

func TestIsPtrThatCanDerefPointerToInt(t *testing.T) {
	myI := 345
	Assert(t).That(&myI, IsPtrThat(Equals(345)))
}

func TestIsPtrThatCanDerefPointerToString(t *testing.T) {
	myI := "Something"
	Assert(t).That(&myI, IsPtrThat(Equals("Something")))
}
