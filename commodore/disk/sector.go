package disk

import (
	"encoding/binary"

	"retroio/storage"
)

// Sector stores the data for each sector of a track.
// The first two bytes store the next Track/Sector locations followed by 254 bytes of data.
type Sector [256]byte

func (s *Sector) Read(reader *storage.Reader) error {
	return binary.Read(reader, binary.LittleEndian, s)
}
