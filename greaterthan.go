package hamcrest

import (
	"fmt"
	"reflect"
)

type Number interface {
}

type IntegerGreaterThanMatcher struct {
	MatchAgainst int64
}

func (me *IntegerGreaterThanMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}
	var toMatch int64

	switch other.(type) {
	case uint, uint8, uint16, uint64, uint32:
		toMatch = int64(reflect.ValueOf(other).Uint())
		if toMatch < 0 {
			result.IsMatched = true
			result.Description = fmt.Sprintf("%d is greater than %d", other, me.MatchAgainst)
			return result
		}
	case int, int8, int16, int64, int32:
		toMatch = reflect.ValueOf(other).Int()
	}

	if toMatch > me.MatchAgainst {
		result.IsMatched = true
		result.Description = fmt.Sprintf("%d is greater than %d", other, me.MatchAgainst)
	} else {
		result.Description = fmt.Sprintf("%d is NOT greater than %d", other, me.MatchAgainst)
	}

	return result
}
func (me *IntegerGreaterThanMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("greater than %d", me.MatchAgainst)
}

type UIntegerGreaterThanMatcher struct {
	MatchAgainst uint64
}

func (me *UIntegerGreaterThanMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}
	var toMatch uint64

	switch other.(type) {
	case uint, uint8, uint16, uint64, uint32:
		toMatch = reflect.ValueOf(other).Uint()
	case int, int8, int16, int64, int32:
		toMatch = uint64(reflect.ValueOf(other).Int())
	}

	if toMatch > me.MatchAgainst {
		result.IsMatched = true
		result.Description = fmt.Sprintf("%d is greater than %d", toMatch, me.MatchAgainst)
	} else {
		result.Description = fmt.Sprintf("%d is NOT greater than %d", toMatch, me.MatchAgainst)
	}

	return result
}
func (me *UIntegerGreaterThanMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("greater than %d", me.MatchAgainst)
}
