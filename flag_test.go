package features_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bavix/features"
)

const (
	Dishing features.Flag = iota
	Washing
	Cooking
	Driving
)

func TestOptions_Has(t *testing.T) {
	toggles := features.New(Dishing, Washing, Cooking)

	require.True(t, toggles.Has(Cooking, Washing, Dishing))
	require.True(t, toggles.Has(Dishing))
	require.True(t, toggles.Has(Cooking))
	require.True(t, toggles.Has(Washing))

	require.False(t, toggles.Has(Cooking, Washing, Dishing, Driving))
	require.False(t, toggles.Has(Driving))
}
