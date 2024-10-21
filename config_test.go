package gogroup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetDebugMode(t *testing.T) {
	SetDebugMode(true)
	require.True(t, debugMode)

	SetDebugMode(false)
	require.False(t, debugMode)
}
