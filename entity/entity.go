package entity

import uuid "github.com/satori/go.uuid"

type Entity interface {
	ID() uuid.UUID
	Name() string
	Health() int
	SetHealth(hp int)
	Step()

	// Damage returns the attack power
	Damage() int
	// Armour returns the armour power
	Armour() int

	AddXP(xp int)
}

func Attack(p1, p2 Entity) {
	for {
		Hit(p1, p2)
		if p2.Health() <= 0 {
			break
		}

		Hit(p2, p1)
		if p1.Health() <= 0 {
			break
		}
	}
}

func Hit(p1, p2 Entity) {
	var damage = -(p2.Armour() - p1.Damage())
	if damage < 0 {
		damage = 0
	}

	p2.SetHealth(p2.Health() - damage)
}
