package hamcrest

import (
	"fmt"
	"reflect"
)

type IsNilMatcher struct{}

func (me *IsNilMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}

	if other == nil || reflect.ValueOf(other).IsNil() {
		result.IsMatched = true
		result.Description = fmt.Sprintf("%s is nil", other)
	} else {
		result.Description = fmt.Sprintf("%s is NOT nil", other)
	}

	return result
}
func (me *IsNilMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("is nil")
}
