package hamcrest_test

import (
	. "github.com/pepinns/go-hamcrest"
	"testing"
)

type TestObject struct {
	FieldOne string
	FieldTwo []int
}

func TestHasFieldMatcherMatchesOnPogoField(t *testing.T) {
	Assert(t).That(&TestObject{FieldOne: "stringy"}, HasFieldThat("FieldOne", Equals("stringy")))
}

// func TestHasFieldMatcherMatchesOnPogoField2(t *testing.T) {
// 	Assert(t).That(&TestObject{FieldTwo: []int{23, 12, 55, 33}}, HasFieldThat(Contains("Two"), AllOf(Contains(12), Contains(522225))))
// }

func TestHasFieldDescriptionIsClear(t *testing.T) {
	AssertFailureMessage(t, &TestObject{FieldOne: "stringyful"}, HasFieldThat("FieldOne", Equals("notmatched")), Equals(`failed to match [
  failed [FieldOne:stringyful] because "stringyful" is not equal to "notmatched"
]`))
}

// type TestNested struct {
// 	InnerPtr *TestObject
// }

// func TestNestedHasFieldWithNulls(t *testing.T) {
// 	tt := &TestNested{}
// 	Assert(t).That(tt, HasField("InnerPtr", HasField("FieldOne", "Four")))
// }

func TestHasFieldDescriptionIsClearWhenComplexMatchersUsed(t *testing.T) {
	AssertFailureMessage(t, &TestObject{FieldTwo: []int{23, 12, 55, 33}}, HasFieldThat(Contains("Two"), AllOf(Contains(12), Contains(55))), Equals(`matched items [
  matched [FieldTwo:<[]int Value>] because "FieldTwo" contains string "Two" and All Matched (
    matched items [
      [1] '12' is equal to '12'
    ]
    matched items [
      [2] '55' is equal to '55'
    ]
  )
]`))
}
