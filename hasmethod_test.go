package hamcrest_test

import (
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

type TestObjectWithGetMethod struct {
	ValueToReturn string
}

func (me *TestObjectWithGetMethod) GetValue() string {
	return me.ValueToReturn
}
func (me *TestObjectWithGetMethod) GetTwoValues() (string, string) {
	return me.ValueToReturn, "TestSecondValue"
}

type TestInterfaceForGetMethod interface {
	GetValue() string
	GetTwoValues() (string, string)
}

func TestHasMethodCanAssertValueOfGetMethods(t *testing.T) {
	var tt TestInterfaceForGetMethod
	tt = &TestObjectWithGetMethod{"TestObjectValue"}
	Assert(t).That(tt, HasMethodThatReturns("GetValue", "TestObjectValue"))
}

func TestHasMethodCanAssertMultipleReturnValuesOfGetMethods(t *testing.T) {
	// var tt TestInterfaceForGetMethod
	tt := &TestObjectWithGetMethod{"TestObjectValue"}
	Assert(t).That(tt, HasMethodThatReturns("GetTwoValues", "TestObjectValue", "TestSecondValue"))
}
func TestHasMethodResponseIsClear(t *testing.T) {
	AssertFailureMessage(t, &TestObjectWithGetMethod{"TestValue"}, HasMethodThatReturns("GetValue", "TestValue"), Contains(`matched [GetValue:<TestValue>] because "GetValue" is equal to "GetValue" and`, `"TestValue" is equal to "TestValue"`))
}
func TestHasMethodResponseIsClearWhenTwoReturnValues(t *testing.T) {
	AssertFailureMessage(t, &TestObjectWithGetMethod{"TestValue"}, HasMethodThatReturns("GetTwoValues", "TestValue", "TestSecondValue"), Contains(`is equal to "TestSecondValue"`))
}

type testVal struct {
	Value string
}
type testValMethod struct {
	val  *testVal
	val2 *testVal
}

func (me *testValMethod) Funcy() (*testVal, *testVal) {
	return me.val, me.val2
}
func TestCanMatchHasFieldAfterHasMethod(t *testing.T) {
	tt := &testValMethod{&testVal{"value"}, &testVal{"value2"}}

	Assert(t).That(tt,
		HasMethodThatReturns("Funcy",
			HasField("Value", "value"),
			HasField("Value", "value2"),
		),
	)
}
