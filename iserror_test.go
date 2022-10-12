package hamcrest_test

import (
	"fmt"
	"testing"

	. "github.com/pepinns/go-hamcrest"
)

type isErrStruct struct {
}

func Test_IsError(t *testing.T) {
	baseError := fmt.Errorf("BaseError")

	t.Run("succeeds using an error", func(t *testing.T) {
		err := fmt.Errorf("error")
		Assert(t).That(err, IsError(err))
	})
	t.Run("succeeds using a wrapped error", func(t *testing.T) {
		err := fmt.Errorf("error: %w", baseError)
		Assert(t).That(err, IsError(baseError))
	})
	t.Run("fails when errors mismatch", func(t *testing.T) {
		err := fmt.Errorf("error1")
		err2 := fmt.Errorf("error2")
		Assert(t).That(err, Not(IsError(err2)))
	})
	t.Run("fails when wrapped errors mismatch", func(t *testing.T) {
		err := fmt.Errorf("error2: %w", fmt.Errorf("error3"))
		Assert(t).That(err, Not(IsError(baseError)))
	})
	t.Run("fails when input is not an error", func(t *testing.T) {
		Assert(t).That(42, Not(IsError(baseError)))
	})
}
