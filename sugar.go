package hamcrest

import (
	"fmt"
	"reflect"
	"regexp"
)

// HasPrefix Match a prefix against a string
func HasPrefix(val string) Matcher {
	return &StringPrefixMatcher{Other: val}
}

// RegexMatches Match against a regular expression
func RegexMatches(expr string) Matcher {
	reg, err := regexp.Compile(expr)
	if err != nil {
		panic(err)
	}
	return &RegexMatcher{reg}
}

// GreaterThan match if greater than
//
// Will compare different numbers types without requiring you to cast the numbers yourself.
func GreaterThan(o interface{}) Matcher {
	switch val := o.(type) {
	case int8, int16, int32, int, int64:
		return &IntegerGreaterThanMatcher{reflect.ValueOf(val).Int()}
	case uint8, uint16, uint32, uint, uint64:
		return &UIntegerGreaterThanMatcher{reflect.ValueOf(val).Uint()}
	}

	panic(fmt.Sprintf("Don't know how to do GreaterThan match on type: %T", o))
}

// Equals
// compares two things to see if they're equal
//
// This method has looser rules around comparing numbers.  For example uint8(2) == int(2) == int32(2).  This is intentional, to allow for an easier time testing using integer contants, avoiding the need to wrap integer constants with casts.
func Equals(o interface{}) Matcher {
	if o == nil {
		return &IsNilMatcher{}
	}
	switch val := o.(type) {
	case string:
		return &StringEqualsMatcher{MatchValue: val}
	case uint32, uint, uint64:
		return &UIntegerEqualsMatcher{MatchValue: reflect.ValueOf(val).Uint()}
	case int32, int, int8, int16, int64:
		return &IntegerEqualsMatcher{MatchValue: reflect.ValueOf(val).Int()}
	case float32, float64:
		return NewFormatEqualsMatcher(o, formatFloat)
	case bool:
		return NewFormatEqualsMatcher(o, formatBool)
	case []byte:
		return NewFormatEqualsMatcher(o, formatBytes)
	default:
		return &DeepEqualsMatcher{MatchValue: o}
	}
	panic(fmt.Sprintf("Don't know how to do equals match on type: %T", o))
}

// Contains
// matches against a string, or a list type.
//
// If the arguments passed to Contains are strings, then this matches if they all are substrings of
// the test value.
// If any of the arguments are not strings, then they are assumed to be Matchers, in this case
// this is treated as a List search, and will match if the test value is an array or a slice.
// Each array/slice item is matched against all of the supplied matchers.
func Contains(o ...interface{}) Matcher {
	if len(o) == 1 {
		return makeContainsMatcher(o[0])
	}

	matchers := make([]interface{}, len(o))
	for idx, val := range o {
		matchers[idx] = makeContainsMatcher(val)
	}
	return AllOf(matchers...)
}
func makeContainsMatcher(o interface{}) Matcher {
	switch val := o.(type) {
	case string:
		return &StringContainsMatcher{MatchValue: formatString(val)}
	}

	return &ListContainsMatcher{ItemMatcher: WrapMatcher(o)}
}

// ContainsSequence
// matches against a specific sequence of items
//
// This can be used to match against Array or Slice types
// each of the matchers must match a contiguous set of indices
// in they order they appear
func ContainsSequence(o ...interface{}) Matcher {
	sc := &SequenceContainsMatcher{}
	for _, m := range o {
		sc.Matchers = append(sc.Matchers, WrapMatcher(m))
	}
	return sc
}

// WrapMatcher
// Wraps the given interface value into a Matcher.
//
// If the value is a Matcher, it is returned, if it is not, then it is
// wrapped in an Equals matcher.
// Input should be Array, Chan, Map, Slice, or String
func WrapMatcher(o interface{}) Matcher {
	if matcher, ok := o.(Matcher); ok {
		return matcher
	}
	return Equals(o)
}

// HasLen
// Compares the input's length against the supplied matcher
func HasLen(o interface{}) Matcher {
	return &HasLenMatcher{WrapMatcher(o)}
}

// WrapAllMatcher  wraps the supplied interface{} values as matchers.
//
// Create a slice of matchers, by passing each input to WrapMatcher() to
// convert to a Matcher.
func WrapAllMatcher(o ...interface{}) []Matcher {
	matchers := make([]Matcher, len(o))
	for idx, val := range o {
		matchers[idx] = WrapMatcher(val)
	}
	return matchers
}

// Not negates the given matcher.
func Not(toNegate Matcher) Matcher {
	return &NotMatcher{toNegate}
}

// HasItem
// Matches on a map structure.
// If the keyMatcher matches a key from the map, then we evaluate the valueMatcher on that key's value.
// To match on just a key, you could use HasKey
func HasItem(keyMatcher interface{}, valueMatcher interface{}) Matcher {
	return &HasItemMatcher{
		KeyMatcher:   WrapMatcher(keyMatcher).(Matcher),
		ValueMatcher: WrapMatcher(valueMatcher).(Matcher),
	}
}

// HasKey
// Matches on a key of a map
// To match a key/value combo use HasItem
func HasKey(keyMatcher interface{}) Matcher {
	return HasItem(keyMatcher, IsAnything())
}

// IsAnything matches anything.
func IsAnything() Matcher {
	return &IsAnythingMatcher{}
}

// IsPtrThat
// de-references the test value, and then passes it to the
// matcher given to this function.  This saves you from having to check for nil, before dereferencing pointers in your test assertions.
func IsPtrThat(o interface{}) Matcher {
	return &IsPointerMatcher{WrapMatcher(o)}
}

// HasField checks that astruct has a field, and then validates the matcher against that struct field's value.
func HasField(fieldNameMatcher interface{}, fieldValueMatcher interface{}) Matcher {
	return HasFieldThat(fieldNameMatcher, fieldValueMatcher)
}

// HasFieldThat checks that astruct has a field, and then validates the matcher against that struct field's value.
func HasFieldThat(fieldNameMatcher interface{}, fieldValueMatcher interface{}) Matcher {
	return &HasFieldMatcher{FieldMatcher: WrapMatcher(fieldNameMatcher), ValueMatcher: WrapMatcher(fieldValueMatcher)}
}

// HasMethodThatReturns
// checks for a method named @methodName, and then attempts to match its
// return values against the supplied matchers, in order.
//
// If the method has 2 return values, you must supply 2 matchers, same for 3, 4 etc.
//
// If you don't care about one of the return values, simply use IsAnything() to match on anything
func HasMethodThatReturns(methodName string, methodResultMatcher ...interface{}) Matcher {
	return &HasMethodMatcher{MethodName: methodName, MethodResultMatchers: WrapAllMatcher(methodResultMatcher...)}
}

// AllOf
// Logical AND of all the supplied matchers
//
// ie.
//
// This will match regardless of the ordering of the slice, since we're doing
// 2 separate checks to see if the item is in the slice.
// Assert(t).That([]string{"string1", "string2", "string4"}, AllOf(
//
//	Contains("string1"),
//	Contains("string2"),
//
// ))
func AllOf(matchers ...interface{}) Matcher {
	tmatchers := make([]Matcher, len(matchers))
	for idx, m := range matchers {
		tmatchers[idx] = m.(Matcher)
	}
	return &AllOfMatcher{Matchers: tmatchers}
}

// AnyOf
// Logical OR of all the supplied matchers
func AnyOf(matchers ...interface{}) Matcher {
	tmatchers := make([]Matcher, len(matchers))
	for idx, m := range matchers {
		tmatchers[idx] = m.(Matcher)
	}
	return &AnyOfMatcher{Matchers: tmatchers}
}

// IsNil matches only nil values
func IsNil() Matcher {
	return &IsNilMatcher{}
}

// IsTrue matches only true values
func IsTrue() Matcher {
	return NewFormatEqualsMatcher(true, formatBool)
}

// IsFalse matches only true values
func IsFalse() Matcher {
	return NewFormatEqualsMatcher(false, formatBool)
}

// IsError using errors.Is to match
func IsError(err error) Matcher {
	return &IsErrorMatcher{err}
}

// ideas ...
// IsSerializedProtobufThat  ... to deserialize a protobuf, and then apply given matchers against it...  could be useful with the field/method matchers

// IsMsgPack...
