package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	var min, max int64 = 100, 1000
	random := RandomInt(min, max)

	require.True(t, random >= min)
	require.True(t, random <= max)
}


