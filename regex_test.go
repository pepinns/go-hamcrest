package hamcrest_test

import (
	. "github.com/pepinns/go-hamcrest"
	"testing"
)

func TestRegexMatcherMatchesRegex(t *testing.T) {
	Assert(t).That("james.test", RegexMatches("^james.*"))
}
func TestRegexMatcherNotMatchesRegex(t *testing.T) {
	Assert(t).That("james.test", Not(RegexMatches("^james$")))
}
