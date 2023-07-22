package sort

import "strings"

type SortDirection int

const (
	Ascending SortDirection = iota
	Descending
)

func GetSortDirection(s string) string {
	switch strings.ToLower(s) {
	case "asc":
		return "ASC"
	case "desc":
		return "DESC"
	default:
		return "ASC"
	}
}
