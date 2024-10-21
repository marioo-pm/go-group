package gogroup

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoGroup(t *testing.T) {
	t.Run("successful_goroutines", func(t *testing.T) {
		group := NewGroup()
		var firstFunctionExecuted, secondFunctionExecuted bool

		group.Go(func() error {
			firstFunctionExecuted = true
			return nil
		})

		group.Go(func() error {
			secondFunctionExecuted = true
			return nil
		})

		require.NoError(t, group.Wait())

		t.Run("make_sure_both_function_are_executed", func(t *testing.T) {
			require.True(t, firstFunctionExecuted)
			require.True(t, secondFunctionExecuted)
		})
	})

	t.Run("failed_goroutine", func(t *testing.T) {
		group := NewGroup()

		group.Go(func() error {
			return errors.New("failed goroutine")
		})

		require.Error(t, group.Wait())
	})

	t.Run("panic_in_goroutine", func(t *testing.T) {
		SetDebugMode(true)
		defer SetDebugMode(false)

		group := NewGroup()

		group.Go(func() error {
			panic("panic in goroutine")
		})

		require.Error(t, group.Wait())
	})
}
