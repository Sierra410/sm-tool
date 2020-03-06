package main

// type iuid uint64

// func (self iuid) Bytes() []byte {
// 	b := [8]byte{}
// 	for i := 0; i < 8; i++ {
// 		b[i] = byte((self >> (0x8 * i)) & 0xFF)
// 	}

// 	return b[:]
// }

// func (self iuid) String() string {
// 	return string(self.Bytes())
// }

// func iuidFromBytes(b []byte) iuid {
// 	if len(b) != 8 {
// 		panic("iuid must be 8 bytes")
// 	}

// 	var id iuid = 0
// 	for i := 0; i < 8; i++ {
// 		id |= iuid(b[i]) << (i * 0x8)
// 	}

// 	return id
// }

// func iuidFromString(s string) iuid {
// 	return iuidFromBytes([]byte(s))
// }

// var iuidCounter = make(chan iuid)

// func init() {
// 	go func() {
// 		for i := iuid(0); ; i++ {
// 			iuidCounter <- i
// 		}
// 	}()
// }
