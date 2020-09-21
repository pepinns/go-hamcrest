package hamcrest

import (
	"fmt"
	"reflect"
)

type HasItemMatcher struct {
	KeyMatcher   Matcher
	ValueMatcher Matcher
}

func (me *HasItemMatcher) Match(other interface{}) MatchResult {
	result := &KeyValueResult{}

	val := reflect.Indirect(reflect.ValueOf(other))
	if val.Kind() == reflect.Map {
		for _, keyVal := range val.MapKeys() {
			itemResult := &KeyResult{KeyDescription: keyVal.String(), ValueDescription: "<>"}
			itemResult.KeyResult = me.KeyMatcher.Match(keyVal)
			if itemResult.KeyResult.Matched() {
				itemVal := val.MapIndex(keyVal)
				itemResult.ValueDescription = itemVal.String()
				if itemVal.CanInterface() {
					itemResult.ValueResult = me.ValueMatcher.Match(itemVal.Interface())
				} else {
					itemResult.ValueResult = &SimpleResult{IsMatched: false, Description: fmt.Sprintf("could not get interface of %s", itemVal)}
				}
			}
			result.Add(itemResult)
			if itemResult.Matched() {
				return result
			}
		}
	}
	return result
}

func (me *HasItemMatcher) WriteDescription(output DescriptionWriter) {
	output.NewLine()
	output.WriteString("has an item with key ")
	output.IncreaseIndent()
	me.KeyMatcher.WriteDescription(output)
	output.DecreaseIndent()
	output.WriteString(" and whose value ")
	output.IncreaseIndent()
	me.ValueMatcher.WriteDescription(output)
	output.DecreaseIndent()
}
