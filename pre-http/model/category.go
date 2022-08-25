package model

import "encoding/json"

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

func (c Category) UnmarshalJSON(data []byte) (err error) {
	var category string
	if err := json.Unmarshal(data, &category); err != nil {
		return err
	}
	return nil
}

func (c Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}
