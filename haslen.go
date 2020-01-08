package hamcrest

import (
	"fmt"
	"reflect"
)

type HasLenMatcher struct {
	LengthMatcher Matcher
}

func (me *HasLenMatcher) Match(other interface{}) MatchResult {

	result := &HasLenResult{LenMatcher: me.LengthMatcher}
	val := reflect.Indirect(reflect.ValueOf(other))
	switch val.Kind() {
	case reflect.Slice,
		reflect.Chan,
		reflect.Array,
		reflect.String:
		result.len = val.Len()
		result.LenResult = me.LengthMatcher.Match(val.Len())
	default:
		lr := &SimpleResult{}
		lr.IsMatched = false
		lr.Description = fmt.Sprintf("%s (type %T) does not have a Length", other, other)
		result.LenResult = lr
	}
	return result
}

func (me *HasLenMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("has length ")
	me.LengthMatcher.WriteDescription(output)
}

type HasLenResult struct {
	len        int
	LenResult  MatchResult
	LenMatcher Matcher
}

func (me *HasLenResult) Matched() bool {
	return me.LenResult.Matched()
}

func (me *HasLenResult) WriteFailureReason(output DescriptionWriter) {
	output.WriteString("length ")
	// output.WriteStringf("length: %d", me.len)
	// if me.Matched() {
	// 	output.WriteString(" was ")
	// } else {
	// 	output.WriteString(" was not ")
	// }
	// me.LenMatcher.WriteDescription(output)
	me.LenResult.WriteFailureReason(output)
}
