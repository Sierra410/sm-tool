package main

import (
	"fmt"
	"math/rand"
)

func validateUuid(s string) bool {
	if len(s) != 36 {
		return false
	}

	for i, r := range s {
		switch i {
		case 8, 13, 18, 23:
			if r != '-' {
				return false
			}
		default:
			if (s[i] < '0' || s[i] > '9') &&
				(s[i] < 'a' || s[i] > 'f') &&
				(s[i] < 'A' || s[i] > 'F') {
				return false
			}
		}
	}

	return true
}

func newUuid4() string {
	return fmt.Sprintf(
		"%08x-%04x-4%04x-%04x-%012x",
		rand.Uint64()%0xFFFFFFFF,
		rand.Uint64()%0xFFFF,
		rand.Uint64()%0xFFF,
		rand.Uint64()%0xFFFF,
		rand.Uint64()%0xFFFFFFFFFFFF,
	)
}
