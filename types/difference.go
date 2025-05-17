package types

import (
	"fmt"
	"strings"
)

type (
	Change struct {
		Description string
		Name        string
		Type        string
		Major       bool
		Minor       bool
	}

	ChangeSet struct {
		Change
		Details []*Change
	}
)

func NewChangeSet(details []*Change) *ChangeSet {
	result := &ChangeSet{
		Details: details,
	}

	majors, minors := 0, 0
	types := make(map[string]int)
	for _, d := range details {
		types[d.Type]++
		if d.Major {
			majors++
		}
		if d.Minor {
			minors++
		}
	}

	if majors > 0 {
		result.Description = fmt.Sprintf("%d breaking changes", majors)
		result.Name = "major revision"
	} else if minors > 0 {
		result.Description = fmt.Sprintf("%d additions", minors)
		result.Name = "minor revision"
	} else {
		result.Description = fmt.Sprintf("%d fixes", len(details))
		result.Name = "bug fix"
	}

	result.Major = majors > 0
	result.Minor = minors > 0

	typeItems := make([]string, 0, len(types))
	for t, c := range types {
		typeItems = append(typeItems, fmt.Sprintf("%s (%d)", t, c))
	}
	result.Type = strings.Join(typeItems, ", ")

	return result
}
