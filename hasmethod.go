package hamcrest

import (
	"fmt"
	"reflect"
	"strings"
)

type HasMethodMatcher struct {
	MethodName           string
	MethodResultMatchers []Matcher
}

func (me *HasMethodMatcher) Match(input interface{}) MatchResult {
	result := &KeyValueResult{}

	inVal := reflect.ValueOf(input)

	meth := inVal.MethodByName(me.MethodName)
	kResult := &KeyResult{KeyDescription: me.MethodName, ValueDescription: "<>"}
	methodMatcher := Equals(me.MethodName)
	kResult.KeyResult = methodMatcher.Match(me.MethodName)

	methodReturn := meth.Call(nil)

	if len(methodReturn) != len(me.MethodResultMatchers) {
		return &SimpleResult{Description: fmt.Sprintf("Method call resulted in %d values, and not %d", len(methodReturn), len(me.MethodResultMatchers))}
	}

	methodResults := &MethodResults{}
	sb := strings.Builder{}
	sb.WriteString("<")
	for idx, matcher := range me.MethodResultMatchers {
		if idx > 0 {
			sb.WriteRune(',')
		}
		var res MatchResult
		if methodReturn[idx].CanInterface() {
			res = matcher.Match(methodReturn[idx].Interface())
		} else {
			res = &SimpleResult{IsMatched: false, Description: fmt.Sprintf("could not get interface of return value %d: %s", idx, methodReturn[idx])}
		}
		methodResults.Results = append(methodResults.Results, res)
		sb.WriteString(methodReturn[idx].String())

	}
	sb.WriteString(">")
	kResult.ValueResult = methodResults

	kResult.ValueDescription = sb.String()
	result.Add(kResult)

	return result
}

type MethodResults struct {
	Results []MatchResult
}

func (me *MethodResults) Matched() bool {
	for _, result := range me.Results {
		if !result.Matched() {
			return false
		}
	}
	return true
}

func (me *MethodResults) WriteFailureReason(writer DescriptionWriter) {
	reset := writer.IncreaseIndent(1)
	defer reset()

	for _, result := range me.Results {
		writer.NewLine()
		result.WriteFailureReason(writer)
	}

}

func (me *HasMethodMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("has a method with a name")
	reset := output.IncreaseIndent(1)
	defer reset()
	output.NewLine()
	methodMatcher := Equals(me.MethodName)
	methodMatcher.WriteDescription(output)
	output.DecreaseIndent(1)
	output.NewLine()
	output.WriteStringf("and when called with no arguments returns %d values", len(me.MethodResultMatchers))
	output.IncreaseIndent(1)
	for i, matcher := range me.MethodResultMatchers {
		output.NewLine()
		output.WriteStringf("[%d] ", i)
		matcher.WriteDescription(output)
	}
}
