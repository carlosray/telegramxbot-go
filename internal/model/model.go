package model

import "fmt"

type HandlePolicy int

const (
	ALL HandlePolicy = iota
	FIRST
)

func (hp HandlePolicy) String() string {
	switch hp {
	case ALL:
		return "ALL"
	case FIRST:
		return "FIRST"
	default:
		return fmt.Sprintf("%d", hp)
	}
}
