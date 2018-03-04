package tape

import "fmt"

// PureData
// ID: 14h (20d)
// This is the same as in the turbo loading data block, except that it has no pilot or sync pulses.
type PureData struct {
	ZeroBitPulse    uint16  // WORD      Length of ZERO bit pulse
	OneBitPulse     uint16  // WORD      Length of ONE bit pulse
	UsedBits        uint8   // BYTE      Used bits in last byte (other bits should be 0) (e.g. if this is 6, then the bits used (x) in the last byte are: xxxxxx00, where MSb is the leftmost bit, LSb is the rightmost bit)
	Pause           uint16  // WORD      Pause after this block (ms.)
	Length          uint16  // N BYTE[3] Length of data that follow
	LengthSpareByte uint8   // NOTE: `length` above uses only 2-bytes for value but specification says 3-bytes, so this is for the spare.
	Data            []uint8 // BYTE[N]   Data as in .TAP files
}

func (p *PureData) Process(file *File) {
	p.OneBitPulse = file.ReadShort()
	p.ZeroBitPulse = file.ReadShort()
	p.UsedBits, _ = file.ReadByte()
	p.Pause = file.ReadShort()
	p.Length = file.ReadShort()
	p.LengthSpareByte, _ = file.ReadByte()

	// Yep, we're discarding the data for the moment
	file.ReadBytes(int(p.Length))
}

func (p PureData) Id() int {
	return 20
}

func (p PureData) Name() string {
	return "Pure Data"
}

// Metadata returns a human readable string of the block data
func (p PureData) Metadata() string {
	str := ""
	str += fmt.Sprintf("OneBitPulse:  %d\n", p.OneBitPulse)
	str += fmt.Sprintf("ZeroBitPulse: %d\n", p.ZeroBitPulse)
	str += fmt.Sprintf("UsedBits:     %d\n", p.UsedBits)
	str += fmt.Sprintf("Pause:        %d\n", p.Pause)
	str += fmt.Sprintf("Length:       %d\n", p.Length)

	return str
}