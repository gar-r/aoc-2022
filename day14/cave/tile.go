package cave

type Tile int

const (
	Brick = iota
	Sand
	Source
)

func (t Tile) String() string {
	switch t {
	case Brick:
		return "#"
	case Sand:
		return "o"
	case Source:
		return "+"
	default:
		return "."
	}
}
