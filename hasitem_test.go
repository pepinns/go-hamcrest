package hamcrest_test

import (
	. "github.com/pepinns/go-hamcrest"
	"testing"
)

func TestHasItemMatchesWhenItemInMap(t *testing.T) {
	mm := make(map[string]string)

	mm["test"] = "value"
	Assert(t).That(mm, HasItem("test", "value"))
}

func Test2HasItemMatchesWhenItemInMap(t *testing.T) {
	mm := make(map[string]string)

	mm["test"] = "value3"
	Assert(t).That(mm, Not(HasItem("test", "value")))
}
func TestHasItemMessageIsClear(t *testing.T) {
	mm := make(map[string]string)
	mm["test"] = "value"

	AssertFailureMessage(t, mm, HasItem("test", "value"), Equals(`matched items [
  matched [test:value] because "test" is equal to "test" and "value" is equal to "value"
]`))
}

func TestHasItemMessageIsClearWhenKeyMatchesButItemDoesNot(t *testing.T) {
	mm := make(map[string]string)
	mm["test"] = "value"

	AssertFailureMessage(t, mm, HasItem("test", "value3"), Equals(`failed to match [
  failed [test:value] because "value" is not equal to "value3"
]`))
}
func TestHasItemMessageIsClearWhenKeyDoesNotMatch(t *testing.T) {
	mm := make(map[string]string)
	mm["test"] = "value"

	AssertFailureMessage(t, mm, HasItem("test1", "value3"), Equals(`failed to match [
  failed [test:<>] because "test" is not equal to "test1"
]`))
}
