package hamcrest_test

import (
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

func TestAnyOfMatchesWhenAnyChildrenMatch(t *testing.T) {
	Assert(t).That(1, AnyOf(Equals(2), Equals(1)))
}
func TestAnyOfNotMatchesWhenAllFailure(t *testing.T) {
	Assert(t).That(1, Not(AnyOf(Equals(10), Equals("foo"), Equals(100))))
}
