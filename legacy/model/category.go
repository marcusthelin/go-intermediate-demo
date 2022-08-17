package model

type Category int

const (
	GROCESSORY Category = iota
	BOOKS
	BABY
	TOYS
)

func (c Category) String() string {
	switch c {
	case GROCESSORY:
		return "GROCESSORY"
	case BABY:
		return "BABY"
	case BOOKS:
		return "BOOKS"
	case TOYS:
		return "TOYS"
	}
	return "unknown"
}
