package hamcrest

import (
	"fmt"
	"reflect"
)

type HasFieldMatcher struct {
	FieldMatcher Matcher
	ValueMatcher Matcher
}

func (me *HasFieldMatcher) Match(input interface{}) MatchResult {
	toMatch := reflect.Indirect(reflect.ValueOf(input))
	if !toMatch.IsValid() {
		return &SimpleResult{
			IsMatched:   false,
			Description: fmt.Sprintf("(%v) was invalid", input)}
	}
	result := &KeyValueResult{}
	foundField := false

	toMatch.FieldByNameFunc(func(name string) bool {
		keyRes := &KeyResult{KeyDescription: name, ValueDescription: "<>"}
		keyRes.KeyResult = me.FieldMatcher.Match(name)
		if keyRes.KeyResult.Matched() {
			foundField = true
			fieldVal := toMatch.FieldByName(name)
			keyRes.ValueDescription = fieldVal.String()
			if fieldVal.CanInterface() {
				keyRes.ValueResult = me.ValueMatcher.Match(fieldVal.Interface())
			} else {
				keyRes.ValueResult = &SimpleResult{IsMatched: false, Description: fmt.Sprintf("could not get interface of %s", fieldVal)}
			}
		}
		result.Add(keyRes)
		if keyRes.Matched() {
			return true
		}
		return false
	})

	if foundField {
		failuresCopy := result.Failures
		result.Failures = make([]*KeyResult, 0)
		//remove all results that don't match the field name
		for _, res := range failuresCopy {
			if res.KeyResult.Matched() {
				result.Failures = append(result.Failures, res)
			}
		}
	}
	return result
}

func (me *HasFieldMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("has a field name that")
	reset := output.IncreaseIndent(1)
	output.NewLine()
	me.FieldMatcher.WriteDescription(output)
	output.DecreaseIndent(1)
	output.NewLine()
	output.WriteString(" and has field value ")
	output.IncreaseIndent(1)
	output.NewLine()
	me.ValueMatcher.WriteDescription(output)
	reset()
}
