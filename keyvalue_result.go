package hamcrest

type KeyResult struct {
	KeyDescription string
	KeyResult      MatchResult

	ValueDescription string
	ValueResult      MatchResult
}

func (me *KeyResult) Matched() bool {
	if me.ValueResult == nil {
		return false
	}
	return me.KeyResult.Matched() && me.ValueResult.Matched()
}
func (me *KeyResult) WriteFailureReason(writer DescriptionWriter) {
	if me.Matched() {
		writer.WriteStringf("matched [%s:%s] because ", me.KeyDescription, me.ValueDescription)
		me.KeyResult.WriteFailureReason(writer)
		writer.WriteString(" and ")
		me.ValueResult.WriteFailureReason(writer)
	} else {
		writer.WriteStringf("failed [%s:%s] because ", me.KeyDescription, me.ValueDescription)
		if !me.KeyResult.Matched() {
			me.KeyResult.WriteFailureReason(writer)
			if me.ValueResult != nil {
				writer.WriteString(" and ")
			}
		}
		if me.ValueResult != nil {
			me.ValueResult.WriteFailureReason(writer)
		}
	}
}

type KeyValueResult struct {
	Failures []*KeyResult
	Sucesses []*KeyResult
}

func (me *KeyValueResult) Matched() bool {
	return len(me.Sucesses) > 0
}
func (me *KeyValueResult) WriteFailureReason(output DescriptionWriter) {
	if me.Matched() {
		output.WriteString("matched items [")
		reset := output.IncreaseIndent(1)
		for _, res := range me.Sucesses {
			output.NewLine()
			res.WriteFailureReason(output)
		}
		reset()
		output.NewLine()
		output.WriteString("]")
	} else {
		output.WriteString("failed to match [")
		reset := output.IncreaseIndent(1)
		for _, res := range me.Failures {
			output.NewLine()
			res.WriteFailureReason(output)
		}
		reset()
		output.NewLine()
		output.WriteString("]")

	}
}

func (me *KeyValueResult) Add(keyResult *KeyResult) {
	if keyResult.Matched() {
		me.Sucesses = append(me.Sucesses, keyResult)
	} else {
		me.Failures = append(me.Failures, keyResult)
	}
}
