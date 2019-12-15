// Package image provides an io.Reader interface, along with various byte reading
// helper functions for processing image data files.
//
// To make this reader easier to use, most of these functions omit the handling
// of errors during the read operation. As we are working with known header/data
// structures, this has been done solely for ease of use.
// Should an EOF ever be reached...well...then there is something seriously wrong
// with the image file, so a `fatal` is triggered.
package storage

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// Image reader, using the bufio.Reader to allow for Peeking.
type Reader struct {
	reader *bufio.Reader
}

// NewReader first converts the regular reader to a buffered reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(r)}
}

func (r Reader) Read(p []byte) (int, error) {
	n, err := io.ReadFull(r.reader, p)
	if n < len(p) {
		fmt.Printf("Read %d bytes, expected %d\n", n, len(p))
	}
	return n, err
}

// ReadByte reads and returns a single byte.
func (r Reader) ReadByte() byte {
	b, _ := r.reader.ReadByte()
	return b
}

// Buffered returns the number of bytes left in the buffer.
func (r Reader) Buffered() int {
	return r.reader.Buffered()
}

// Size returns the size of the reader.
func (r Reader) Size() int {
	return r.reader.Size()
}

// ReadNextBytes reads a variable length of bytes from the reader.
// WARNING: if you try reading past the EOF this func will panic.
func (r Reader) ReadNextBytes(number int) []byte {
	b := make([]byte, number)

	n, err := r.Read(b)
	if err != nil || n < number {
		log.Fatal(fmt.Errorf("ReadNextBytes failed spectacularly!\n%s", err))
	}

	return b
}

// ReadLong reads a value from the reader, then converts
// the little endian ordered bytes to a uint32.
func (r Reader) ReadLong() uint32 {
	b := r.ReadNextBytes(4)
	return binary.LittleEndian.Uint32(b[:])
}

// ReadShort reads a value from the reader, then converts
// the little endian ordered bytes to a uint16.
func (r Reader) ReadShort() uint16 {
	b := r.ReadNextBytes(2)
	return binary.LittleEndian.Uint16(b[:])
}

// ReadShortToSigned reads a value from the reader, then
// converts to a `signed` int.
func (r Reader) ReadShortToSigned() int16 {
	b := r.ReadNextBytes(2)
	return int16(binary.LittleEndian.Uint16(b[:]))
}

// Peek returns the next n bytes without advancing the reader.
func (r Reader) Peek(n int) ([]byte, error) {
	b, err := r.reader.Peek(n)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PeekShort reads two bytes without advancing the reader, then
// converts the little endian ordered bytes to an uint16.
func (r Reader) PeekShort() (uint16, error) {
	b, err := r.reader.Peek(2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b[:]), nil
}

// Discard wraps the bufio.Discard function
func (r Reader) Discard(n int) (int, error) {
	return r.reader.Discard(n)
}

// BytesToLong converts a slice of 4 little endian ordered bytes.
func (r Reader) BytesToLong(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b[:])
}
