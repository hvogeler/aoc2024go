package maze

import "fmt"

type Reindeer struct {
	id       int
	heading  HeadingType
	position Position
	score    int
	state    ReindeerState
}

func (r Reindeer) String() string {
	prefix := "Dead "
	if r.state == alive {
		prefix = "Alive"
	}
	s := fmt.Sprintf("%s Reeindeer %d at %s heading %s has score of %d\n", prefix, r.id, r.position, r.heading, r.score)
	return s
}

func (r *Reindeer) Kill() {
	r.state = dead
}

type ReindeerState int

const (
	alive ReindeerState = iota
	dead
)
