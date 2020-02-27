package hamcrest_test

import (
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

func TestContainsSequenceMatchesIfStringArrayContainsSequence(t *testing.T) {
	ary := []string{"one", "two", "three"}
	Assert(t).That(ary, ContainsSequence("two", "three"))
	Assert(t).That(ary, Not(ContainsSequence("two", "four", "three")))
}

// func TestContainsSequenceErrorMessageIsClear(t *testing.T) {
// 	ary := []string{"one", "two", "three"}
// 	AssertFailureMessage(t, ary,
// 		ContainsSequence("two", "four", "three"),
// 		Equals(`"Fred"`))
// }
