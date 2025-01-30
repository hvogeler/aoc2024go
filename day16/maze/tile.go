package maze

type Tile interface {
	TileType() TileType
	String() string
}

// ***** Wall *****

type Wall struct{}

func (w Wall) TileType() TileType {
	return WallType
}

func (w Wall) String() string {
	return string(w.TileType())
}


// ***** Unused *****

type Unused struct{}

func (u Unused) TileType() TileType {
	return UnusedType
}

func (u Unused) String() string {
	return string(u.TileType())
}


// ***** Start *****

type Start struct{}

func (s Start) TileType() TileType {
	return StartType
}

func (w Start) String() string {
	return string(w.TileType())
}

// ***** Finish *****

type Finish struct{}

func (f Finish) TileType() TileType {
	return FinishType
}

func (w Finish) String() string {
	return string(w.TileType())
}

type TileType string

const (
	WallType   TileType = "#"
	UnusedType TileType = "."
	StartType  TileType = "S"
	FinishType TileType = "E"
	TrackMarkType  TileType = "T"
)
