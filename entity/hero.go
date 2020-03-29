package entity

import "math/rand"

const defaultHeroAge = 18
const defaultHeroHealth = 25

// Hero describes hero entity
type Hero struct {
	id   int64
	name string
	x, y int

	age    int
	nage   int // number of months before your birthday
	health int
}

func (h *Hero) X() int {
	return h.x
}

func (h *Hero) Y() int {
	return h.y
}

func (h *Hero) SetX(x int) {
	h.x = x
}

func (h *Hero) SetY(y int) {
	h.y = y
}

// NewHero creates new hero with default params
func NewHero(name string) *Hero {
	return &Hero{
		id:     nextEntityID.Inc(),
		name:   name,
		age:    defaultHeroAge,
		health: defaultHeroHealth,
	}
}

// ID returns hero's ID
func (h Hero) ID() int64 {
	return h.id
}

// Name returns hero's name
func (h Hero) Name() string {
	return h.name
}

// Age returns hero's age
func (h Hero) Age() int {
	return h.age
}

// Health returns hero's health
func (h Hero) Health() int {
	return h.health
}

func (h *Hero) Step() {
	h.nage += 1
	if h.nage >= 365 {
		h.nage = 0
		h.age++
	}

	pos := rand.Intn(4)
	switch pos {
	case 0:
		Up(h)
	case 1:
		Down(h)
	case 2:
		Left(h)
	case 3:
		Right(h)
	}
}
