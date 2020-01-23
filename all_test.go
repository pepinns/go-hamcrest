package hamcrest_test

import (
	. "github.com/pepinns/go-hamcrest"
	"testing"
)

func TestAllOfMatchesWhenAllChildrenMatch(t *testing.T) {
	Assert(t).That(1, AllOf(Equals(1), Equals(1)))
}
func TestAllOfNotMatchesWhenOneFailure(t *testing.T) {
	Assert(t).That(1, Not(AllOf(Equals(1), Equals("foo"), Equals(1))))
}
