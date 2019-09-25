package hamcrest

type IsAnythingMatcher struct{}

func (me *IsAnythingMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}
	result.IsMatched = true
	result.Description = "is anything"
	return result
}

func (me *IsAnythingMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("matches")
}
