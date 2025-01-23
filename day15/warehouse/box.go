package warehouse

const BoxStr1 string = "O"
const BoxStr2 string = "[]"

type Box struct {
	position Location
}

func NewBox(x, y int) Box {
	return Box{position: NewLocation(x, y)}
}

func (box Box) GpsCoord() int {
	return (box.position.y)*100 + box.position.x
}
