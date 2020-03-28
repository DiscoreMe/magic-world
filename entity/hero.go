package entity

const defaultHeroAge = 18
const defaultHeroHealth = 25

type Hero struct {
	id     int64
	name   string
	age    int
	nage   int // number of months before your birthday
	health int
	x, y   int
}

func NewHero(name string) *Hero {
	return &Hero{
		id:     nextEntityID(),
		name:   name,
		age:    defaultHeroAge,
		health: defaultHeroHealth,
	}
}

// ID gets character ID
func (c Hero) ID() int64 {
	return c.id
}

// Name gets name ID Hero
func (c Hero) Name() string {
	return c.name
}

// Age gets age Hero
func (c Hero) Age() int {
	return c.age
}

func (c Hero) Health() int {
	return c.health
}

func (c *Hero) Step() {
	c.nage += 1
	if c.nage >= 12 {
		c.nage = 0
		c.age++
	}
}

func (c *Hero) SetPos(x, y int) {
	c.x, c.y = x, y
}

func (c Hero) Pos() (int, int) {
	return c.x, c.y
}
