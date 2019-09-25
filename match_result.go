package hamcrest

type SimpleResult struct {
	Description string
	IsMatched   bool
}

func (me *SimpleResult) Matched() bool {
	return me.IsMatched
}
func (me *SimpleResult) WriteFailureReason(output DescriptionWriter) {
	output.WriteString(me.Description)
}
