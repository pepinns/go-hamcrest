package hamcrest

import "reflect"

type HasFieldMatcher struct {
	FieldMatcher Matcher
	ValueMatcher Matcher
}

func (me *HasFieldMatcher) Match(input interface{}) MatchResult {

	toMatch := reflect.Indirect(reflect.ValueOf(input))

	result := &KeyValueResult{}
	toMatch.FieldByNameFunc(func(name string) bool {
		keyRes := &KeyResult{KeyDescription: name, ValueDescription: "<>"}
		keyRes.KeyResult = me.FieldMatcher.Match(name)
		if keyRes.KeyResult.Matched() {
			fieldVal := toMatch.FieldByName(name)
			keyRes.ValueDescription = fieldVal.String()
			if fieldVal.CanInterface() {
				keyRes.ValueResult = me.ValueMatcher.Match(fieldVal.Interface())
			} else {
				keyRes.ValueResult = me.ValueMatcher.Match(fieldVal)
			}
		}
		result.Add(keyRes)
		if keyRes.Matched() {
			return true
		}
		return false
	})
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
