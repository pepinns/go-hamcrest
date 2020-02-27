package hamcrest

import (
	"fmt"
	"regexp"
)

type RegexMatcher struct {
	MatchAgainst *regexp.Regexp
}

func (me *RegexMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}

	toMatch := fmt.Sprintf("%s", other)

	if me.MatchAgainst.Match([]byte(toMatch)) {
		result.IsMatched = true
		result.Description = fmt.Sprintf("%s matches %s", toMatch, me.MatchAgainst)
	} else {
		result.Description = fmt.Sprintf("%s does NOT match %s", toMatch, me.MatchAgainst)
	}

	return result
}
func (me *RegexMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("matches regex %s", me.MatchAgainst.String())
}
