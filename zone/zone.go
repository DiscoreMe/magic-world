package zone

const (
	ZoneTypeCity = iota
)

type Zone struct {
	Type   int
	Name   string
	Course int
}

func NewZone(ztype int, name string, course int) *Zone {
	return &Zone{
		Type:   ztype,
		Name:   name,
		Course: course,
	}
}
