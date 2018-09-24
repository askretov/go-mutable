package mutable

import (
	"strconv"
)

type Status int

const (
	NotChanged Status = iota
	Removed
	Added
	Changed
)

// String implements Stringer interface for Status
func (m Status) String() string {
	switch m {
	case NotChanged:
		return "NotChanged"
	case Removed:
		return "Removed"
	case Added:
		return "Added"
	case Changed:
		return "Changed"
	default:
		return strconv.Itoa(int(m))
	}
}
