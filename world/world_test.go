package world_test

import (
	"testing"

	"github.com/DiscoreMe/magic-world/entity"
	"github.com/DiscoreMe/magic-world/world"
	"github.com/stretchr/testify/assert"
)

const worldWidth, worldHeight = 10, 10

func TestWorld(t *testing.T) {
	hero := entity.NewHero("Test")

	w := world.NewWorld(worldWidth, worldHeight)
	w.AddEntity(worldWidth/2, worldHeight/2, hero)

	x, y := hero.X(), hero.Y()
	days := w.Days()

	w.Step()

	assert.NotEqual(t, x, hero.X())
	assert.NotEqual(t, y, hero.Y())
	assert.Equal(t, days+1, w.Days())
	assert.NoError(t, w.ExportToFile("test.world"))
}
