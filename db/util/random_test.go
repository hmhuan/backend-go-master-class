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

func TestRandomString(t *testing.T) {
	length := 10
	random := RandomString(length)

	require.Equal(t, length, len(random))
}

func TestRandomOwner(t *testing.T) {
	owner1 := RandomOwner()

	require.Equal(t, 10, len(owner1))
}

func TestRandomBalance(t *testing.T) {
	balance := RandomBalance()

	require.True(t, balance >= 100)
	require.True(t, balance <= 1000)
}

func TestRandomCurrency(t *testing.T) {
	currencies := []string{"USD", "JPY", "VND", "CAD", "EUR"}
	random := RandomCurrency()

	require.Contains(t, currencies, random)
}