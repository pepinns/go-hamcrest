Go Hamcrest 
=========================

This library is a port of the [Hamcrest java library](http://hamcrest.org/).

Its purpose is to make tests easier to develop, as well as making test failures easier to understand.

It makes your tests easier to write by creating a declarative language for introspecting values to eliminate the need for writing logic in tests to check the results of an operation.

Another helpful side effect of using this is avoiding having to check pointers or interface types for nil before checking their properties.

You eliminate all the code thats not performing useful testing work, which reduces the maintenance overhead, and also makes modifying tests and keeping them up to date much easier.


Examples
-------------
    

The very simplest example is comparing two strings.

The following code

``` go
package hamcrest_test

import (
    . "hamcrest"
    "testing"
)

func TestStringEqualsFailsWithAsserter(t *testing.T) {
    Assert(t).That("prefixFAIL", Equals("prefix"))
}
```

Produces the following test output

    --- FAIL: TestStringEqualsFailsWithAsserter (0.00s)
        assert.go:68: failed asserting that "prefixFAIL" 
              equal to "prefix"
                  because
              "prefixFAIL" is not equal to "prefix"
              
              

Here is a very contrived example of a complex matcher.

The following Code

``` go
import (
    . "hamcrest"
    "testing"
)

type TestObject struct {
    FieldOne string
    FieldTwo []int
}

func TestHasFieldMatcherMatchesOnPogoField2(t *testing.T) {
    Assert(t).That(&TestObject{FieldTwo: []int{23, 12, 55, 33}}, HasFieldThat(Contains("Two"), AllOf(Contains(12), Contains(522225))))
}
```

And that test failure produces the following output.   Notice the detail it gives you.

    --- FAIL: TestHasFieldMatcherMatchesOnPogoField2 (0.00s)
        assert.go:68: failed asserting that &hamcrest_test.TestObject{FieldOne:"", FieldTwo:[]int{23, 12, 55, 33}} 
              has a field name that
                contains string "Two"
              and has field value 
                All of the following:(  contains a list item that equal to 12 AND  contains a list item that equal to 522225 ) 
                  because
              failed to match [
                failed [FieldOne:<>] because "FieldOne" does not contain string "Two"
                failed [FieldTwo:<[]int Value>] because All Of (failed to match [
                  [0] '23' is not equal to '522225'
                  [1] '12' is not equal to '522225'
                  [2] '55' is not equal to '522225'
                  [3] '33' is not equal to '522225'
                ])
              ]
    FAIL


This is obviously very contrived, but you can see the power of this matching library in this example.  Note that in order to perform this without such a library you'd have to write the code
to use reflection to find the proper field on this object, after which you'd have to iterate the slice of integers and compare on each to figure out if your values were in the slice.  

Instead, you can just write a few functions.


Contributing Matchers
-----------------------------------

All are encouraged to contribute new matchers to this.  Its fairly easy to add a new one, there are only a couple of requirements.

There are 2 important interfaces.  [Matcher](assert.go "hamcrest.Matcher") and [MatchResult](assert.go "hamcrest.MatchResult").

``` go
type Matcher interface {
	Match(interface{}) MatchResult
	WriteDescription(DescriptionWriter)
}

type MatchResult interface {
	Matched() bool
	WriteFailureReason(DescriptionWriter)
}
```

The Matcher simply needs to provide 2 methods.  

1 for describing itself and what its going to match on. `WriteDescription(DescriptionWriter)`

This can be fairly simple, or complex, depending on what you're doing.  If you're matching on things that are brown, your description might be "is brown". To indicate that you're going to try and match on things that are brown.

The actual matching logic takes place in the `Match(interface{}) MatchResult` method.  

Here is where you'd perform your tests against the input data, and collect any information about the failure that you'd like to report via the MatchResult.  In the simplest of cases, you can get away with using the [SimpleResult](match_result.go "hamcrest.SimpleResult").
The `SimpleResult` simply allows you to set a Description field, with your failure description.  


If you have anything more complicated to say than a simple string, then its time to write your own Result structure.  One example of this is the [KeyValueResult for HasItem/HasField](keyvalue_result.go), since it needs to track which keys/values it tested in order to give as much info as possible in the output.

In this result, I've created some compound structures to simplify the logic of figuring out which things were not matching and printing them out in a clear way.

If you look back at the git history, you'll see that my result structure also makes the logic of figuring out if we matched alot simpler.


Open/Closed
---------------------

The interfaces and structures of this library are designed to be open/closed, that is Open to change, but Closed for modification.  This means that the interfaces and utilities that you need to make your own matchers interact with this libary are all exposed.  You should be able to build something that meets the hamcrest.Matcher interface in your own project/package and it will work just fine.

This should make it easier to write matchers for your own project, that maybe don't fit in a general purpose library.  If you have some logic that is specific to your application or test setup, you can write your own matchers.


Abstracting Matchers
----------------------------

A common pattern with this type of matching setup is to create helper functions that wrap up a useful matcher or set of matchers.

An example of this is for matching an `http.Header` on an http.Response or `http.Request` structure.

``` go
func HttpHeader(name interface{}, value interface{}) {
    return HasFieldThat("Header", HasItem(name, Contains(value)))
}
```

Instead of making a custom `AssertXXXX`  Function, which only works on that one exact set of inputs, and cannot be combined into a larger set of matchers.

Having all your matching combined into one Assertion means that whenever your test fails, you'll get a message that clearly shows you all the things that were wrong, and not just the first bit of the test to trip.  Often that information is enough for you to realize that your change to seemingly unrelated code was actually the culprit in breaking these tests.  You no longer need to write out long unhelpful strings to your `t.Fatalf** calls which tend to get stale over time, and sometimes even lie to you because the test was updated, but the assertion failure message maybe wasn't.

TODO
----------------------

There are a number of things that would make this more useful.  In no specific order

    * Refactor the DescriptionWriter to allow a structured format, like json, so that we could better integrate test failures with reporting tools.
    * Add more matchers
      ** Make new sub-packages for matching protocol buffers from []byte slices, so you don't have to Unmarshal them yourself.
      ** Same goes for other structured formats like json,yaml,msgpack, etc...
      ** More Equality matchers... right now we're relying on `reflect.DeepEquals`, but with more coding we could get specific equality checkers for Maps and Lists, that would make the errors a bit easier on the eyes.
