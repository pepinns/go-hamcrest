package hamcrest

import (
	"errors"
	"fmt"
)

type IsErrorMatcher struct {
	Other error
}

func (me *IsErrorMatcher) Match(input interface{}) MatchResult {
	result := &SimpleResult{}
	if to_check, ok := input.(error); ok {
		if errors.Is(to_check, me.Other) {
			result.IsMatched = true
			result.Description = fmt.Sprintf("\"%s\" is \"%s\"", to_check, me.Other)
			return result
		}
		result.Description = fmt.Sprintf("\"%s\" is not \"%s\"", to_check, me.Other)
	} else {
		result.Description = fmt.Sprintf("%+v is %T and not an error", input, input)
	}
	return result
}

func (me *IsErrorMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("an error that matches ")
	output.WriteString(me.Other.Error())
}
