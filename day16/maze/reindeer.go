package maze

import "fmt"

type Reindeer struct {
	id      int
	heading HeadingType
	score   int
	state   ReindeerState
	visited []Position
	reachedEnd bool
}

func NewReindeer(id int, position Position) Reindeer {
	return Reindeer{
		id:      id,
		heading: Right,
		visited: []Position{position},
	}
}

func (r Reindeer) AlreadyVisited(checkPosition Position) bool {
	for _, p := range r.visited {
		if checkPosition == p {
			return true
		}
	}
	return false
}

func (r Reindeer) IsAlive() bool {
	return r.state == alive
}

func (r Reindeer) Position() Position {
	return r.visited[len(r.visited)-1]
}

func (r *Reindeer) SetPosition(p Position) {
	r.visited = append(r.visited, p)
}

func (r Reindeer) String() string {
	prefix := "Dead "
	if r.state == alive {
		prefix = "Alive"
	}
	s := fmt.Sprintf("%s Reindeer %d at %s heading %s has score of %d, tracklen(%d)\n", prefix, r.id, r.Position(), r.heading, r.score, len(r.visited))
	return s
}

func (r Reindeer) Clone(newId int, newHeading HeadingType, score int) Reindeer {
	newReindeer := Reindeer{
		id:      newId,
		heading: newHeading,
		score:   score,
	}
	newReindeer.visited = make([]Position, len(r.visited))
	copy(newReindeer.visited, r.visited)
	return newReindeer
}

func (r *Reindeer) Kill(reason string) {
	r.state = dead
	if reason[:3] == "End" {
		fmt.Printf("Reindeer %d killed: %s, Score %d\n", r.id, reason, r.score)
	}
}

type ReindeerState int

const (
	alive ReindeerState = iota
	dead
)
