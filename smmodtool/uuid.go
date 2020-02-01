package main

import (
	"fmt"
	"math/rand"
)

// 0-7    309e1830
// 8      -
// 9-12   7865
// 13     -
// 14-17  421b
// 18     -
// 19-22  8886
// 23     -
// 24-35  c47f2ce0010e

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
		case 14:
			if r != '4' {
				return false
			}
		case 19:
			switch r {
			case '8', '9', 'a', 'A', 'b', 'B':
			default:
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

func randomUuid() string {
	return fmt.Sprintf(
		"%08x-%04x-4%03x-%04x-%012x",
		rand.Uint64()%0xFFFFFFFF,
		rand.Uint64()%0xFFFF,
		rand.Uint64()%0xFFF,
		rand.Uint64()%0xFFFF&0x3FFF|0x8000,
		rand.Uint64()%0xFFFFFFFFFFFF,
	)
}
