package entity

const defaultHeroAge = 18
const defaultHeroHealth = 25

// Hero describes hero entity
type Hero struct {
	id     int64
	name   string
	health int

	age  int
	nage int // number of months before your birthday
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

// Health returns hero's health
func (h Hero) Health() int {
	return h.health
}

// Age returns hero's age
func (h Hero) Age() int {
	return h.age
}

func (h *Hero) Step() {

}
