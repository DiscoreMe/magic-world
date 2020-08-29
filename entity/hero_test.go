package entity_test

import (
	"github.com/DiscoreMe/magic-world/entity"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

import "testing"

var hero *entity.Hero
var heroName = gofakeit.Name()

func init() {
	hero = entity.NewHero(heroName, 10)
}

func TestHero_Name(t *testing.T) {
	assert.Equal(t, hero.Name(), heroName)
}

func TestHero_ID(t *testing.T) {
	assert.NotEqual(t, hero.ID(), 0)
}

func TestHero_Age(t *testing.T) {
	assert.NotEqual(t, hero.Age(), 0)
}

func TestHero_Health(t *testing.T) {
	assert.NotEqual(t, hero.Health(), 0)
}
