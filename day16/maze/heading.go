package maze

type HeadingType string

const (
	Up    HeadingType = "^"
	Down  HeadingType = "v"
	Left  HeadingType = "<"
	Right HeadingType = ">"
)

func HeadingTypes() []HeadingType {
	return []HeadingType{Up, Down, Left, Right}
}

func (dt HeadingType) Score(newDt HeadingType) int {
	switch dt.TurnRate(newDt) {
	case 0:
		return 1
	case 90:
		return 1000
	case 180:
		return 1000 + 1000
	default:
		panic("Switch exhausted")
	}
}

func (dt HeadingType) TurnRate(newDt HeadingType) int {
	if (dt == Up && newDt == Down) || (dt == Down && newDt == Up) {
		return 180
	}
	if (dt == Right && newDt == Left) || (dt == Left && newDt == Right) {
		return 180
	}
	if (dt == Up || dt == Down) && (newDt == Left || newDt == Right) {
		return 90
	}
	if (newDt == Up || newDt == Down) && (dt == Left || dt == Right) {
		return 90
	}
	if dt == newDt {
		return 0
	}
	panic("Exhausted Ifs")
}
