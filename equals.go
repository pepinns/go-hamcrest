package hamcrest

import (
	"fmt"
	"reflect"
)

type StringEqualsMatcher struct {
	MatchValue string
}

func (me *StringEqualsMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}
	if formatString(other) == me.MatchValue {
		result.IsMatched = true
		result.Description = fmt.Sprintf("\"%s\" is equal to \"%s\"", other, me.MatchValue)
	} else {
		// if strings.Contains(me.MatchValue, "\n") {
		//   result.Description = fmt.Sprintf("\"%s\" is not equal to \"%s\"", other, me.MatchValue) + "\n" + diff.LineDiff(formatString(other), me.MatchValue)
		// } else {
		result.Description = fmt.Sprintf("\"%s\" is not equal to \"%s\"", other, me.MatchValue)
		// }
	}

	return result
}
func (me *StringEqualsMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("equal to \"%s\"", me.MatchValue)
}

type IntegerEqualsMatcher struct {
	MatchValue int64
}

func (me *IntegerEqualsMatcher) Match(other interface{}) MatchResult {
	var toMatch int64
	result := &SimpleResult{}
	switch other.(type) {
	case reflect.Value:
		vv := reflect.Indirect(other.(reflect.Value))
		switch vv.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			toMatch = int64(vv.Uint())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			toMatch = vv.Int()
		}
	//TODO: fix this
	// this is backwards... our MatchValue is what was typed in... the 'other' here represents the value we're asserting..
	// that means that the value we're asserting could in fact overflow int64, so we need to test for that and fail if it does.
	case uint, uint8, uint16, uint64, uint32:
		toMatch = int64(reflect.ValueOf(other).Uint())
	case int, int8, int16, int64, int32:
		toMatch = reflect.ValueOf(other).Int()
	}
	if toMatch == me.MatchValue {
		result.IsMatched = true
		result.Description = fmt.Sprintf("'%d' is equal to '%d'", other, me.MatchValue)
	} else {
		result.Description = fmt.Sprintf("'%d' is not equal to '%d'", other, me.MatchValue)
	}
	return result
}
func (me *IntegerEqualsMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("equal to ")
	output.WriteString(fmt.Sprintf("%d", me.MatchValue))
}

type UIntegerEqualsMatcher struct {
	MatchValue uint64
}

func (me *UIntegerEqualsMatcher) Match(other interface{}) MatchResult {
	result := &SimpleResult{}
	var toMatch uint64

	switch other.(type) {
	case int, int8, int16, int64, int32:
		to_match := reflect.ValueOf(other).Int()
		if to_match < 0 {
			result.IsMatched = false
			result.Description = fmt.Sprintf("'%d' is < 0 and cannot be converted to uint", to_match)
			return result
		}
		toMatch = uint64(to_match)
	case uint, uint8, uint16, uint64, uint32:
		toMatch = reflect.ValueOf(other).Uint()
	}
	if toMatch == me.MatchValue {
		result.IsMatched = true
		result.Description = fmt.Sprintf("'%d' is equal to '%d'", other, me.MatchValue)
	} else {
		result.Description = fmt.Sprintf("'%d' is not equal to '%d'", other, me.MatchValue)
	}

	return result
}
func (me *UIntegerEqualsMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteString("equal to ")
	output.WriteString(fmt.Sprintf("%d", me.MatchValue))
}

func formatString(b interface{}) string {
	return fmt.Sprintf("%s", b)
}
func formatBytes(b interface{}) string {
	return fmt.Sprintf("%x", b)
}
func formatFloat(b interface{}) string {
	return fmt.Sprintf("%.4f", b)
}
func formatBool(b interface{}) string {
	return fmt.Sprintf("%t", b)
}
func NewFormatEqualsMatcher(matchValue interface{}, format func(interface{}) string) Matcher {
	fm := &FormatEqualsMatcher{}
	fm.MatchValue = format(matchValue)
	fm.Format = format
	return fm
}

type FormatEqualsMatcher struct {
	StringEqualsMatcher
	Format func(interface{}) string
}

func (me *FormatEqualsMatcher) Match(other interface{}) MatchResult {
	return me.StringEqualsMatcher.Match(me.Format(other))
}

type DeepEqualsMatcher struct {
	MatchValue interface{}
}

func (me *DeepEqualsMatcher) Match(input interface{}) MatchResult {
	result := &SimpleResult{}
	if reflect.DeepEqual(input, me.MatchValue) {
		result.IsMatched = true
		result.Description = fmt.Sprintf("\"%#v\" is equal to \"%#v\"", me.MatchValue, input)
	} else {
		result.Description = fmt.Sprintf("\"%#v\" is not equal to \"%#v\"", me.MatchValue, input)
	}
	return result
}

func (me *DeepEqualsMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("is equal to \"%#v\"", me.MatchValue)
}
