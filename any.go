package hamcrest

type AnyOfMatcher struct {
	Matchers []Matcher
}

func (me *AnyOfMatcher) Match(other interface{}) MatchResult {
	result := &AnyOfResult{}

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
func (me *AnyOfMatcher) WriteDescription(writer DescriptionWriter) {
	writer.WriteString(" Any of the following:( ")
	writer.NewLine()
	for idx, matcher := range me.Matchers {
		if idx > 0 {
			writer.NewLine()
			writer.IncreaseIndent(2)
			writer.WriteString(" - OR - ")
			writer.DecreaseIndent(2)
			writer.NewLine()
		}
		matcher.WriteDescription(writer)
	}
	writer.WriteString(" ) ")
}

type AnyOfResult struct {
	Failures []MatchResult
	Sucesses []MatchResult
}

func (me *AnyOfResult) Matched() bool {
	return len(me.Sucesses) > 0
}
func (me *AnyOfResult) WriteFailureReason(output DescriptionWriter) {
	if me.Matched() {
		output.WriteString("Matched (")
		reset := output.IncreaseIndent(1)
		for _, item := range me.Sucesses {
			output.NewLine()
			item.WriteFailureReason(output)
		}
		reset()
		output.NewLine()
	} else {
		output.WriteString("None Of (")
		output.IncreaseIndent(1)
		for _, item := range me.Failures {
			output.NewLine()
			item.WriteFailureReason(output)
		}
		output.DecreaseIndent(1)
	}
	output.WriteString(")")
}
