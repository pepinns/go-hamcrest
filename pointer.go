package hamcrest

import (
	"reflect"
)

type IsPointerMatcher struct {
	ValueMatcher Matcher
}

func (me *IsPointerMatcher) Match(input interface{}) MatchResult {
	val := reflect.ValueOf(input)
	iVal := reflect.Indirect(val)
	if iVal.CanInterface() {
		return me.ValueMatcher.Match(iVal.Interface())
	}
	return me.ValueMatcher.Match(iVal)
	// }
	// result := &SimpleResult{}
	// result.IsMatched = false
	// result.Description = fmt.Sprintf("%#v was not a pointer", input)
	// return result
}

func (me *IsPointerMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("is a pointer whose value")
	reset := output.IncreaseIndent(1)
	defer reset()
	output.NewLine()
	me.ValueMatcher.WriteDescription(output)
}
