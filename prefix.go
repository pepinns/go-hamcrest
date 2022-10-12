package hamcrest

import (
	"fmt"
	"strings"
)

type StringPrefixMatcher struct {
	Other string
}

func (me *StringPrefixMatcher) Match(input interface{}) MatchResult {
	result := &SimpleResult{}
	if to_check, ok := input.(string); ok {
		if strings.HasPrefix(to_check, me.Other) {
			result.IsMatched = true
			result.Description = fmt.Sprintf("\"%s\" starts with \"%s\"", to_check, me.Other)
			return result
		}
		result.Description = fmt.Sprintf("\"%s\" does not start with \"%s\"", to_check, me.Other)
	} else {
		result.Description = fmt.Sprintf("%+v is %T and not a string", input, input)
	}
	return result
}

func (me *StringPrefixMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("a string that starts with ")
	output.WriteString(me.Other)
}
