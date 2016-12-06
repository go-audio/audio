package audio

import (
	"encoding/binary"
	"errors"
)

var (
	ErrInvalidBuffer = errors.New("invalid buffer")
)

// Format is a high level representation of the underlying data.
type Format struct {
	// NumChannels is the number of channels contained in the data
	NumChannels int
	// SampleRate is the sampling rate in Hz
	SampleRate int
	// BitDepth is the number of bits of data for each sample
	BitDepth int
	// Endianess indicate how the byte order of underlying bytes
	Endianness binary.ByteOrder
}

type Buffer interface {
	PCMFormat() *Format
	Clone() Buffer
	NumFrames() int
	AsFloatBuffer() *FloatBuffer
	AsIntBuffer() *IntBuffer
}

type FloatBuffer struct {
	Format *Format
	Data   []float64
}

func (buf *FloatBuffer) PCMFormat() *Format          { return buf.Format }
func (buf *FloatBuffer) AsFloatBuffer() *FloatBuffer { return buf }

type IntBuffer struct {
	Format *Format
	Data   []int
}

func (buf *IntBuffer) PCMFormat() *Format { return buf.Format }
func (buf *IntBuffer) AsFloatBuffer() *FloatBuffer {
	return nil
}
