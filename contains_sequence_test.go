package hamcrest_test

import (
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

func TestContainsSequenceMatchesIfStringArrayContainsSequence(t *testing.T) {
	ary := []string{"one", "two", "three"}
	Assert(t).That(ary, ContainsSequence("two", "three"))
	Assert(t).That(ary, ContainsSequence("one", "two", "three"))
	Assert(t).That(ary, ContainsSequence("one"))
	Assert(t).That(ary, ContainsSequence("two"))
	Assert(t).That(ary, ContainsSequence("three"))
	Assert(t).That(ary, Not(ContainsSequence("two", "four", "three")))
	Assert(t).That(ary, Not(ContainsSequence("two", "three", "four")))
	Assert(t).That(ary, Not(ContainsSequence("four", "two", "three", "four")))
	Assert(t).That(ary, Not(ContainsSequence("four", "two", "three")))
	Assert(t).That(ary, Not(ContainsSequence("three", "four")))
	Assert(t).That(ary, Not(ContainsSequence("four")))
}

// func TestContainsSequenceErrorMessageIsClear(t *testing.T) {
// 	ary := []string{"one", "two", "three"}
// 	AssertFailureMessage(t, ary,
// 		ContainsSequence("two", "four", "three"),
// 		Equals(`"Fred"`))
// }
