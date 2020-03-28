package world_test

import (
	"github.com/DiscoreMe/magic-world/world"
	"github.com/stretchr/testify/assert"
	"testing"
)

const worldWidth, worldHeight = 10, 10

func TestWorld(t *testing.T) {
	w := world.NewWorld(worldWidth, worldHeight)

	args := []struct {
		X, Y     int
		V        string
		wantType int
	}{
		{1, 1, "test 1 1", 1},
		{2, 2, "test 2 2", 1},
		{2, 3, "test 2 3", 1},
		{2, 4, "test 2 4", 1},
		{4, 2, "test 4 2", 1},
		{worldWidth + 1, worldHeight + 1, "", 0},
	}

	w.CreateLand()

	for _, arg := range args {
		w.Zone.SetMeta(arg.X, arg.Y, arg.V)
		assert.Equal(t, arg.V, w.Zone.Meta(arg.X, arg.Y))
		assert.Equal(t, arg.wantType, w.Zone.Type(arg.X, arg.Y))
	}

	assert.NoError(t, w.ExportToJSON("test.world"))
}
