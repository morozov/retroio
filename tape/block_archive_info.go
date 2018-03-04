package tape

import "fmt"

// ArchiveInfo
// ID: 32h (50d)
// Use this block at the beginning of the tape to identify the title of the game, author,
// publisher, year of publication, price (including the currency), type of software (arcade
// adventure, puzzle, word processor, ...), protection scheme it uses (Speedlock 1, Alkatraz,
// ...) and its origin (Original, Budget re-release, ...), etc. This block is built in a way
// that allows easy future expansion. The block consists of a series of text strings. Each text
// has its identification number (which tells us what the text means) and then the ASCII text.
// To make it possible to skip this block, if needed, the length of the whole block is at the
// beginning of it.
// If all texts on the tape are in English language then you don't have to supply the 'Language' field.
type ArchiveInfo struct {
	Length      uint16 // WORD  Length of the whole block (without these two bytes)
	StringCount uint8  // N BYTE  Number of text strings
	Strings     []Text // TEXT[N] List of text strings
}

type Text struct {
	TypeID uint8 // BYTE  Text identification byte:
	//                       00 - Full title
	//                       01 - Software house/publisher
	//                       02 - Author(s)
	//                       03 - Year of publication
	//                       04 - Language
	//                       05 - Game/utility type
	//                       06 - Price
	//                       07 - Protection scheme/loader
	//                       08 - Origin
	//                       FF - Comment(s)
	Length     uint8  // L BYTE  Length of text string
	Characters []byte // CHAR[L] Text string in ASCII format
}

var categories = map[uint8]string{
	0:   "Title",
	1:   "Publisher",
	2:   "Authors",
	3:   "Year",
	4:   "Language",
	5:   "Category",
	6:   "Price",
	7:   "Loader",
	8:   "Origin",
	255: "Comment",
}

func (a *ArchiveInfo) Process(file *File) {
	a.Length = file.ReadShort()
	a.StringCount, _ = file.ReadByte()

	for i := 0; i < int(a.StringCount); i++ {
		var t Text
		t.TypeID, _ = file.ReadByte()
		t.Length, _ = file.ReadByte()
		for _, c := range file.ReadBytes(int(t.Length)) {
			t.Characters = append(t.Characters, c)
		}
		a.Strings = append(a.Strings, t)
	}
}

func (a ArchiveInfo) Id() int {
	return 50
}

func (a ArchiveInfo) Name() string {
	return "Archive Info"
}

// Metadata returns a human readable string of the block data
func (a ArchiveInfo) Metadata() string {
	str := ""
	for _, b := range a.Strings {
		str += fmt.Sprintf("- %-9s: %s\n", categories[b.TypeID], b.Characters)
	}

	return str
}
