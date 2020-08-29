package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	entity1 Entity
	entity2 Entity
)

func init() {
	entity1 = &Hero{
		health:     10,
		baseDamage: 10,
		baseArmour: 10,
	}
	entity2 = &Hero{
		health:     5,
		baseDamage: 10,
		baseArmour: 1,
	}
}

func TestHit(t *testing.T) {
	var tempHealth = entity2.Health()
	Hit(entity1, entity2)
	assert.LessOrEqual(t, entity2.Health(), tempHealth)
}

func TestAttack(t *testing.T) {
	Attack(entity1, entity2)
	assert.LessOrEqual(t, entity2.Health(), 0)
}
