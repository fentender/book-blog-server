package swagger

import (
	"encoding/binary"
)

// itob returns an 8-byte big endian representation of v.
func itob(v int32) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
