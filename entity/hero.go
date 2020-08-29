package entity

import (
	"github.com/satori/go.uuid"
	"math/rand"
)

const (
	defaultHeroAge    = 18
	defaultHeroHealth = 25
	defaultHeroDamage = 1
	defaultHeroArmour = 1
)

// Hero describes hero entity
type Hero struct {
	id     uuid.UUID
	name   string
	health int

	level int
	xp    int

	age  int
	nage int // number of months before your birthday

	baseDamage int
	baseArmour int
}

// NewHero creates new hero with default params
func NewHero(name string, nage int) *Hero {
	return &Hero{
		id:         uuid.NewV4(),
		name:       name,
		health:     defaultHeroHealth,
		age:        defaultHeroAge,
		nage:       rand.Intn(nage),
		baseDamage: defaultHeroDamage,
		baseArmour: defaultHeroArmour,
	}
}

// ID returns hero's ID
func (h Hero) ID() uuid.UUID {
	return h.id
}

// Name returns hero's name
func (h Hero) Name() string {
	return h.name
}

// Health returns hero's health
func (h Hero) Health() int {
	return h.health
}

func (h *Hero) SetHealth(xp int) {
	h.health = xp
}

// Age returns hero's age
func (h Hero) Age() int {
	return h.age
}

func (h *Hero) Step() {

}

func (h *Hero) Damage() int {
	return h.baseDamage
}

func (h *Hero) Armour() int {
	return h.baseArmour
}

func (h *Hero) AddXP(xp int) {
	h.xp += xp
}
