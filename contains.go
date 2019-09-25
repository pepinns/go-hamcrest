package hamcrest

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
)

type StringContainsMatcher struct {
	MatchValue string
}

func (me *StringContainsMatcher) Match(other interface{}) MatchResult {
	if slicey, ok := other.([]string); ok {
		result := &ListContainsResult{}
		eq := Equals(me.MatchValue)
		for idx, val := range slicey {
			result.Add(idx, val, eq.Match(val))
		}
		return result
	}

	result := &SimpleResult{}
	if reader, ok := other.(io.Reader); ok {
		readBytes, err := ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		other = string(readBytes)
	}

	if strings.Contains(fmt.Sprintf("%s", other), me.MatchValue) {
		result.IsMatched = true
		result.Description = fmt.Sprintf("\"%s\" contains string \"%s\"", other, me.MatchValue)
	} else {
		result.Description = fmt.Sprintf("\"%s\" does not contain string \"%s\"", other, me.MatchValue)
	}
	return result
}
func (me *StringContainsMatcher) WriteDescription(output DescriptionWriter) {
	output.WriteStringf("contains string \"%s\"", me.MatchValue)
}

type ListContainsMatcher struct {
	ItemMatcher Matcher
}

func (me *ListContainsMatcher) Match(other interface{}) MatchResult {
	result := &ListContainsResult{}

	rval := reflect.ValueOf(other)
	switch rval.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rval.Len(); i++ {
			val := rval.Index(i)
			result.Add(i, val.String(), me.ItemMatcher.Match(val))
		}
	}
	return result
}
func (me *ListContainsMatcher) WriteDescription(builder DescriptionWriter) {
	builder.WriteString(" contains a list item that ")
	me.ItemMatcher.WriteDescription(builder)
}

type IndexResult struct {
	MatchResult
	Index           int
	ItemDescription string
}

func (me *IndexResult) WriteFailureReason(output DescriptionWriter) {
	output.WriteStringf("[%d] ", me.Index)
	me.MatchResult.WriteFailureReason(output)

}

type ListContainsResult struct {
	Failures  []*IndexResult
	Successes []*IndexResult
}

func (me *ListContainsResult) Matched() bool {
	return len(me.Successes) > 0
}
func (me *ListContainsResult) WriteFailureReason(output DescriptionWriter) {
	if me.Matched() {
		output.WriteString("matched items [")
		reset := output.IncreaseIndent(1)
		for _, keyRes := range me.Successes {
			output.NewLine()
			keyRes.WriteFailureReason(output)
		}
		reset()
		output.NewLine()
		output.WriteString("]")

	} else {
		output.WriteString("failed to match [")
		reset := output.IncreaseIndent(1)
		for _, keyRes := range me.Failures {
			output.NewLine()
			keyRes.WriteFailureReason(output)
		}
		reset()
		output.NewLine()
		output.WriteString("]")

	}
}

func (me *ListContainsResult) Add(idx int, itemDescription string, result MatchResult) {
	idxResult := &IndexResult{result, idx, itemDescription}
	if idxResult.Matched() {
		me.Successes = append(me.Successes, idxResult)
	} else {
		me.Failures = append(me.Failures, idxResult)
	}
}
