package hamcrest_test

import (
	. "hamcrest"
	"strings"
)

func AssertFailureMessage(t TestT, value interface{}, matcher Matcher, failureMessageMatcher Matcher) {
	t.Helper()
	result := matcher.Match(value)
	description := &FormattingDescriptionWriter{}
	result.WriteFailureReason(description)
	Assert(t).That(description.String(), failureMessageMatcher)
}
func AssertFailureString(t TestT, value interface{}, matcher Matcher, failureMessage string) {
	t.Helper()
	result := matcher.Match(value)
	description := &FormattingDescriptionWriter{}
	result.WriteFailureReason(description)
	Assert(t).That(formatString(description.String()), Equals(formatString(failureMessage)))
}

func formatString(input string) string {
	input = "\n" + input
	input = strings.Replace(input, " ", ".", -1)
	return input
}
