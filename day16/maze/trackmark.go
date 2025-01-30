package maze

type TrackMark struct {
	reindeerId int
	direction DirectionType
	score     int
}

func (tm TrackMark) TileType() TileType {
	return TrackMarkType
}

func (tm TrackMark) String() string {
	return string(tm.TileType())
}

type DirectionType string

const (
	Up    DirectionType = "^"
	Down  DirectionType = "v"
	Left  DirectionType = ">"
	Right DirectionType = "<"
)

func DirectionTypes() []DirectionType {
	return []DirectionType{Up, Down, Left, Right}
}