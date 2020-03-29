package entity

import "math/rand"

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
	if c.nage >= 365 {
		c.nage = 0
		c.age++
	}

	c.randMove()
}

func (c *Hero) randMove() {
	pos := rand.Intn(4)
	switch pos {
	case 0:
		c.Up()
	case 1:
		c.Down()
	case 2:
		c.Left()
	case 3:
		c.Right()
	}
}

func (c *Hero) SetPos(x, y int) {
	c.x, c.y = x, y
}

func (c Hero) Pos() (int, int) {
	return c.x, c.y
}

func (c *Hero) Up() {
	c.y -= 1
}
func (c *Hero) Down() {
	c.y += 1
}
func (c *Hero) Left() {
	c.x -= 1
}
func (c *Hero) Right() {
	c.x += 1
}

type HeroAround struct {
	UpX, UpY       int
	DownX, DownY   int
	LeftX, LeftY   int
	RightX, RightY int
}

func (c *Hero) Around() HeroAround {
	return HeroAround{
		UpX:    c.x,
		UpY:    c.y + 1,
		DownX:  c.x,
		DownY:  c.y - 1,
		LeftX:  c.x - 1,
		LeftY:  c.y,
		RightX: c.x + 1,
		RightY: c.y,
	}
}
