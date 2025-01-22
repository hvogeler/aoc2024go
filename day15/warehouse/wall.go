package warehouse

const WallRune string = "#"

type Wall struct {
	position Location
}

func NewWall(x, y int) Wall {
	return Wall{position: NewLocation(x, y)}
}
