package utils

import (
	"strings"
)

// Bool returns a pointer to the bool value passed in
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}

	return false
}

func UInt64(v uint64) *uint64 {
	return &v
}

func Contains(source []string, input string) bool {
	for _, s := range source {
		if strings.EqualFold(strings.ToLower(s), strings.ToLower(input)) {
			return true
		}
	}

	return false
}
