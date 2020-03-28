package character

var lastCharacterID int64 = 0

type Character struct {
	id   int64
	name string
	age  int
}

// ID gets character ID
func (c Character) ID() int64 {
	return c.id
}

// Name gets name ID
func (c Character) Name() string {
	return c.name
}

// Age gets age Character
func (c Character) Age() int {
	return c.age
}

func NewCharacter() *Character {
	lastCharacterID++
	return &Character{
		id: lastCharacterID,
	}
}