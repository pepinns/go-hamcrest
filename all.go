package hamcrest

type AllOfMatcher struct {
	Matchers []Matcher
}

func (me *AllOfMatcher) Match(other interface{}) MatchResult {
	result := &AllOfResult{}

	for _, matcher := range me.Matchers {
		mResult := matcher.Match(other)
		if !mResult.Matched() {
			result.Failures = append(result.Failures, mResult)
		} else {
			result.Sucesses = append(result.Sucesses, mResult)
		}
	}

	return result
}
func (me *AllOfMatcher) WriteDescription(writer DescriptionWriter) {
	writer.WriteString(" All of the following:( ")
	writer.NewLine()
	for idx, matcher := range me.Matchers {
		if idx > 0 {
			writer.NewLine()
			writer.IncreaseIndent(2)
			writer.WriteString(" - AND - ")
			writer.DecreaseIndent(2)
			writer.NewLine()
		}
		matcher.WriteDescription(writer)
	}
	writer.WriteString(" ) ")
}

type AllOfResult struct {
	Failures []MatchResult
	Sucesses []MatchResult
}

func (me *AllOfResult) Matched() bool {
	return len(me.Failures) == 0
}
func (me *AllOfResult) WriteFailureReason(output DescriptionWriter) {
	if me.Matched() {
		output.WriteString("All Matched (")
		reset := output.IncreaseIndent(1)
		for _, item := range me.Sucesses {
			output.NewLine()
			item.WriteFailureReason(output)
		}
		reset()
		output.NewLine()
	} else {
		output.WriteString("All Of (")
		output.IncreaseIndent(1)
		for _, item := range me.Failures {
			output.NewLine()
			item.WriteFailureReason(output)
		}
		output.DecreaseIndent(1)
	}
	output.WriteString(")")

}
