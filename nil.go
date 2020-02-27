package hamcrest

import (
	"fmt"
	"reflect"
)

type IsNilMatcher struct{}

func (me *IsNilMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}

	if me.matches(other) {
		result.IsMatched = true
		result.Description = fmt.Sprintf("%s is nil", other)
	} else {
		result.Description = fmt.Sprintf("%s is NOT nil", other)
	}

	return result
}

func (me *IsNilMatcher) matches(other interface{}) bool {
	if other == nil {
		return true
	}

	rVal := reflect.ValueOf(other)
	switch rVal.Kind() {
	case reflect.Ptr, reflect.UnsafePointer,
		reflect.Slice, reflect.Interface:
		return rVal.IsNil()
	case reflect.Struct:
		return false
	}

	return false
}
func (me *IsNilMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("is nil")
}
