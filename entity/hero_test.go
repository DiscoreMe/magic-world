package entity_test

import (
	"github.com/DiscoreMe/magic-world/entity"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

import "testing"

var hero *entity.Hero
var heroName = gofakeit.Name()

func TestNewHero(t *testing.T) {
	hero = entity.NewHero(heroName)
}

func TestHero_Name(t *testing.T) {
	assert.Equal(t, hero.Name(), heroName)
}

func TestHero_Step(t *testing.T) {
	age := hero.Age()
	hero.Step()
	assert.Equal(t, age, hero.Age())
	for i := 0; i < 12; i++ {
		hero.Step()
	}
	assert.True(t, age < hero.Age())
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
