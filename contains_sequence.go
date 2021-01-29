package hamcrest

import (
	"fmt"
	"reflect"
)

type SequenceContainsMatcher struct {
	Matchers []Matcher
}

func (me *SequenceContainsMatcher) Match(other interface{}) MatchResult {
	result := &SequenceContainsResult{}

	seqIndex := 0
	rval := reflect.ValueOf(other)
	switch rval.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rval.Len() && seqIndex < len(me.Matchers); i++ {
			val := rval.Index(i)
			nextMatcher := me.Matchers[seqIndex]
			var res MatchResult
			if val.CanInterface() {
				res = nextMatcher.Match(val.Interface())
			} else {
				res = nextMatcher.Match(val)
			}
			if !res.Matched() {
				if seqIndex == 0 {
					continue // we haven't found a match yet
				}
			}
			result.Add(i, seqIndex, res)
			seqIndex++
		}
		if seqIndex < len(me.Matchers) {
			res := &SimpleResult{
				Description: fmt.Sprintf("Ran out of fields to match on. %d matchers left in sequence", len(me.Matchers)-seqIndex),
				IsMatched:   false,
			}
			result.Add(rval.Len(), seqIndex, res)
		}
	}
	return result
}
func (me *SequenceContainsMatcher) WriteDescription(builder DescriptionWriter) {
	builder.WriteString(" contains the following sequence: ")
	for idx, item := range me.Matchers {
		builder.NewLine()
		builder.WriteStringf("[n+%d] ", idx)
		item.WriteDescription(builder)
	}
}

type SeqIndexResult struct {
	MatchResult
	SourceIndex   int
	SequenceIndex int
}

func (me *SeqIndexResult) WriteFailureReason(output DescriptionWriter) {
	indicator := "-"
	if me.MatchResult.Matched() {
		indicator = "+"
	}
	output.WriteStringf("%s[%d:n+%d] ", indicator, me.SourceIndex, me.SequenceIndex)
	me.MatchResult.WriteFailureReason(output)
}

type SequenceContainsResult struct {
	MatchResults []*SeqIndexResult
}

func (me *SequenceContainsResult) Matched() bool {
	for _, res := range me.MatchResults {
		if !res.Matched() {
			return false
		}
	}
	return true
}
func (me *SequenceContainsResult) WriteFailureReason(output DescriptionWriter) {
	output.WriteString("items [")
	reset := output.IncreaseIndent(1)
	for _, keyRes := range me.MatchResults {
		output.NewLine()
		keyRes.WriteFailureReason(output)
	}
	reset()
	output.NewLine()
	output.WriteString("]")

}

func (me *SequenceContainsResult) Add(idx, seqIdx int, result MatchResult) {
	idxResult := &SeqIndexResult{result, idx, seqIdx}
	me.MatchResults = append(me.MatchResults, idxResult)
}
