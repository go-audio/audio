package audio

import (
	"encoding/binary"
	"errors"
)

var (
	// ErrInvalidBuffer is a generic error returned when trying to read/write to an invalid buffer.
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

// Buffer is the representation of an audio buffer.
type Buffer interface {
	// PCMFormat is the format of buffer (describing the buffer content/format).
	PCMFormat() *Format
	// NumFrames returns the number of frames contained in the buffer.
	NumFrames() int
	// AsFloatBuffer returns a float buffer from this buffer.
	AsFloatBuffer() *FloatBuffer
	// AsIntBuffer returns an int buffer from this buffer.
	AsIntBuffer() *IntBuffer
	// Clone creates a clean clone that can be modified without
	// changing the source buffer.
	Clone() Buffer
}

// FloatBuffer is an audio buffer with its PCM data formatted as float64.
type FloatBuffer struct {
	// Format is the representation of the underlying data format
	Format *Format
	// Data is the buffer PCM data as floats
	Data []float64
}

// PCMFormat returns the buffer format information.
func (buf *FloatBuffer) PCMFormat() *Format { return buf.Format }

// AsFloatBuffer implements the Buffer interface and returns itself.
func (buf *FloatBuffer) AsFloatBuffer() *FloatBuffer { return buf }

// AsIntBuffer returns a copy of this buffer but with data truncated to Ints.
func (buf *FloatBuffer) AsIntBuffer() *IntBuffer {
	newB := &IntBuffer{}
	newB.Data = make([]int, len(buf.Data))
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = int(buf.Data[i])
	}
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
		BitDepth:    buf.Format.BitDepth,
		Endianness:  buf.Format.Endianness,
	}
	return newB
}

// Clone creates a clean clone that can be modified without
// changing the source buffer.
func (buf *FloatBuffer) Clone() Buffer {
	if buf == nil {
		return nil
	}
	newB := &FloatBuffer{}
	newB.Data = make([]float64, len(buf.Data))
	copy(newB.Data, buf.Data)
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
		BitDepth:    buf.Format.BitDepth,
		Endianness:  buf.Format.Endianness,
	}
	return newB
}

// NumFrames returns the number of frames contained in the buffer.
func (buf *FloatBuffer) NumFrames() int {
	if buf == nil || buf.Format == nil {
		return 0
	}
	numChannels := buf.Format.NumChannels
	if numChannels == 0 {
		numChannels = 1
	}

	return len(buf.Data) / numChannels
}

// IntBuffer is an audio buffer with its PCM data formatted as int.
type IntBuffer struct {
	// Format is the representation of the underlying data format
	Format *Format
	// Data is the buffer PCM data as ints
	Data []int
}

// PCMFormat returns the buffer format information.
func (buf *IntBuffer) PCMFormat() *Format { return buf.Format }

// AsFloatBuffer returns a copy of this buffer but with data converted to floats.
func (buf *IntBuffer) AsFloatBuffer() *FloatBuffer {
	newB := &FloatBuffer{}
	newB.Data = make([]float64, len(buf.Data))
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = float64(buf.Data[i])
	}
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
		BitDepth:    buf.Format.BitDepth,
		Endianness:  buf.Format.Endianness,
	}
	return newB
}

// AsIntBuffer implements the Buffer interface and returns itself.
func (buf *IntBuffer) AsIntBuffer() *IntBuffer { return buf }

// NumFrames returns the number of frames contained in the buffer.
func (buf *IntBuffer) NumFrame() int {
	if buf == nil || buf.Format == nil {
		return 0
	}
	numChannels := buf.Format.NumChannels
	if numChannels == 0 {
		numChannels = 1
	}

	return len(buf.Data) / numChannels
}
