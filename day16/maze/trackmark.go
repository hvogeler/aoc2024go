package maze

type TrackMark struct {
	reindeerId int
	heading    HeadingType
	score      int
}

func (tm TrackMark) TileType() TileType {
	return TrackMarkType
}

func (tm TrackMark) String() string {
	return string(tm.heading)
}
