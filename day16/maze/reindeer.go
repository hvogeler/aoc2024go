package maze

import "fmt"

type Reindeer struct {
	id int
	heading DirectionType
	position Position
}

func (r Reindeer) String() string {
	return fmt.Sprintf("Reeindeer %d at %s heading %s\n", r.id, r.position, r.heading)
}