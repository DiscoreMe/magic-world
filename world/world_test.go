package world_test

import (
	"github.com/DiscoreMe/magic-world/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcZone(t *testing.T) {
	w := world.NewWorld()

	args := []struct{
		X, Y int
		V string
	}{
		{1,1, "test 1 1"},
		{2,2, "test 2 2"},
		{2,3, "test 2 3"},
		{2,4, "test 2 4"},
		{4,2, "test 4 2"},
	}

	for _, arg := range args {
		w.Zone.SetMeta(arg.X, arg.Y, arg.V)
		assert.Equal(t, arg.V, w.Zone.Meta(arg.X, arg.Y))
	}
}