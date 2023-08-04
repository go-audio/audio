package audio

import "math"

var _ Buffer = (*IntBuffer)(nil)

// IntBuffer is an audio buffer with its PCM data formatted as int.
type IntBuffer struct {
	// Format is the representation of the underlying data format
	Format *Format
	// Data is the buffer PCM data as ints
	Data []int
	// SourceBitDepth helps us know if the source was encoded on
	// 8, 16, 24, 32, 64 bits.
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

// GetSourceBitDepth returns buf.SourceBitDepth if populated, otherwise returns an estimate
// of the source bit depth based on the range of integer values contained in the buffer.
func (buf *IntBuffer) GetSourceBitDepth() int {
	if buf.SourceBitDepth != 0 {
		return buf.SourceBitDepth
	}

	max := int64(0)
	min := int64(0)
	for _, s := range buf.Data {
		if int64(s) > max {
			max = int64(s)
		} else if int64(s) < min {
			min = int64(s)
		}
	}
	if -min > max {
		max = -min
	}

	// 8-bit PCM uses unsigned ints (bytes)
	// Require max > 0 (vals in a silent 8-bit buffer should be ~128)
	if min >= 0 && max > 0 && max <= 255 {
		return 8
	}
	for _, n := range []int{16, 24, 32} {
		// max abs val of an n-bit signed int is 2^(n-1)
		if max <= 1<<(n-1) {
			return n
		}
	}
	return 64
}

// AsFloat32Buffer returns a copy of this buffer but with data converted to float 32.
func (buf *IntBuffer) AsFloat32Buffer() *Float32Buffer {
	newB := &Float32Buffer{}
	newB.Data = make([]float32, len(buf.Data))
	newB.SourceBitDepth = buf.GetSourceBitDepth()
	var toFloat func(int) float32
	if newB.SourceBitDepth == 8 {
		// 8-bit uses unsigned ints
		toFloat = func(d int) float32 {
			return float32(d)/255*2 - 1
		}
	} else {
		factor := 1.0 / math.Pow(2, float64(newB.SourceBitDepth-1))
		toFloat = func(d int) float32 {
			return float32(float64(d) * factor)
		}
	}
	for i := 0; i < len(buf.Data); i++ {
		newB.Data[i] = toFloat(buf.Data[i])
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
