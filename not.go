package hamcrest

type NotMatcher struct {
	ToNegate Matcher
}

func (me *NotMatcher) Match(other interface{}) MatchResult {
	return &NotResult{me.ToNegate.Match(other)}
}
func (me *NotMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("is not ")
	me.ToNegate.WriteDescription(output)
}

type NotResult struct {
	Inner MatchResult
}

func (me *NotResult) Matched() bool {
	return !me.Inner.Matched()
}
func (me *NotResult) WriteFailureReason(output DescriptionWriter) {
	output.WriteString("(")
	reset := output.IncreaseIndent(1)
	output.NewLine()
	me.Inner.WriteFailureReason(output)
	reset()
	output.NewLine()
	output.WriteString(")")
}
