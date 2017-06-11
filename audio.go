package audio

import (
	"errors"
	"math"
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
}

// Buffer is the representation of an audio buffer.
type Buffer interface {
	// PCMFormat is the format of buffer (describing the buffer content/format).
	PCMFormat() *Format
	// NumFrames returns the number of frames contained in the buffer.
	NumFrames() int
	// AsFloatBuffer returns a float 64 buffer from this buffer.
	AsFloatBuffer() *FloatBuffer
	// AsFloat32Buffer returns a float 32 buffer from this buffer.
	AsFloat32Buffer() *Float32Buffer
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

// AsFloat32Buffer implements the Buffer interface and returns a float 32 version of itself.
func (buf *FloatBuffer) AsFloat32Buffer() *Float32Buffer {
	newB := &Float32Buffer{}
	newB.Data = make([]float32, len(buf.Data))
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = float32(buf.Data[i])
	}
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
	}
	return newB
}

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

// Float32Buffer is an audio buffer with its PCM data formatted as float32.
type Float32Buffer struct {
	// Format is the representation of the underlying data format
	Format *Format
	// Data is the buffer PCM data as floats
	Data []float32
}

// PCMFormat returns the buffer format information.
func (buf *Float32Buffer) PCMFormat() *Format { return buf.Format }

// AsFloatBuffer implements the Buffer interface and returns a float64 version of itself.
func (buf *Float32Buffer) AsFloatBuffer() *FloatBuffer {
	newB := &FloatBuffer{}
	newB.Data = make([]float64, len(buf.Data))
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = float64(buf.Data[i])
	}
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
	}
	return newB
}

// AsFloat32Buffer implements the Buffer interface and returns itself.
func (buf *Float32Buffer) AsFloat32Buffer() *Float32Buffer { return buf }

// AsIntBuffer returns a copy of this buffer but with data truncated to Ints.
func (buf *Float32Buffer) AsIntBuffer() *IntBuffer {
	newB := &IntBuffer{}
	newB.Data = make([]int, len(buf.Data))
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = int(buf.Data[i])
	}
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
	}
	return newB
}

// Clone creates a clean clone that can be modified without
// changing the source buffer.
func (buf *Float32Buffer) Clone() Buffer {
	if buf == nil {
		return nil
	}
	newB := &Float32Buffer{}
	newB.Data = make([]float32, len(buf.Data))
	copy(newB.Data, buf.Data)
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
	}
	return newB
}

// NumFrames returns the number of frames contained in the buffer.
func (buf *Float32Buffer) NumFrames() int {
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
	// SourceBitDepth helps us know if the source was encoded on
	// 1 (int8), 2 (int16), 3(int24), 4(int32), 8(int64) bytes.
	SourceBitDepth int
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
	}
	return newB
}

// AsFloat32Buffer returns a copy of this buffer but with data converted to float 32.
func (buf *IntBuffer) AsFloat32Buffer() *Float32Buffer {
	newB := &Float32Buffer{}
	newB.Data = make([]float32, len(buf.Data))
	max := 0
	bitDepth := buf.SourceBitDepth
	// try to guess the bit depths without knowing the source
	if bitDepth == 0 {
		for _, s := range buf.Data {
			if s > max {
				max = s
			}
		}
		bitDepth = 8
		if max > 127 {
			bitDepth = 16
		}
		// greater than int16, expecting int24
		if max > 32767 {
			bitDepth = 24
		}
		// int 32
		if max > 8388607 {
			bitDepth = 32
		}
		// int 64
		if max > 4294967295 {
			bitDepth = 64
		}
	}
	factor := math.Pow(2, 8*float64(bitDepth/8)-1)
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = float32(float64(int64(buf.Data[i])) / factor)
	}
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
	}
	return newB
}

// AsIntBuffer implements the Buffer interface and returns itself.
func (buf *IntBuffer) AsIntBuffer() *IntBuffer { return buf }

// NumFrames returns the number of frames contained in the buffer.
func (buf *IntBuffer) NumFrames() int {
	if buf == nil || buf.Format == nil {
		return 0
	}
	numChannels := buf.Format.NumChannels
	if numChannels == 0 {
		numChannels = 1
	}

	return len(buf.Data) / numChannels
}

// Clone creates a clean clone that can be modified without
// changing the source buffer.
func (buf *IntBuffer) Clone() Buffer {
	if buf == nil {
		return nil
	}
	newB := &IntBuffer{}
	newB.Data = make([]int, len(buf.Data))
	copy(newB.Data, buf.Data)
	newB.Format = &Format{
		NumChannels: buf.Format.NumChannels,
		SampleRate:  buf.Format.SampleRate,
	}
	return newB
}
