package hamcrest_test

import (
	. "hamcrest"
	"testing"
)

func TestStringPrefixReturnsFalseOnDiffString(t *testing.T) {
	eq := &StringPrefixMatcher{"other"}

	result := eq.Match("something different")

	if result.Matched() {
		t.Fatal("expected 2 strings to NOT match")
	}

}
func TestStringPrefixReturnsTrueOnSameString(t *testing.T) {
	eq := &StringPrefixMatcher{"prefix"}

	result := eq.Match("prefix of string")

	if !result.Matched() {
		t.Fatal("expected 'prefix' is prefix of 'prefix of string'")
	}

}

func TestStringPrefixWithAsserter(t *testing.T) {
	Assert(t).That("prefix", HasPrefix("prefix"))
	Assert(t).That("prefix", HasPrefix("pr"))
}
func TestStringPrefixMatcherSuccessReasonIsClear(t *testing.T) {
	AssertFailureMessage(t, "string to match", HasPrefix("string to"), Equals(`"string to match" starts with "string to"`))
}

func TestStringPrefixMatcherFailureReasonIsClear(t *testing.T) {
	AssertFailureMessage(t, "string to match", HasPrefix("failed match"), Equals(`"string to match" does not start with "failed match"`))
}

func TestStringPrefixMatcherFailureReasonIsClearWhenItemIsNotString(t *testing.T) {
	AssertFailureMessage(t, 34, HasPrefix("failed match"), Equals(`34 is int and not a string`))
}
