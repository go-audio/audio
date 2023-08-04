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

// AsFloat32Buffer returns a copy of this buffer but with data converted to float 32.
func (buf *IntBuffer) AsFloat32Buffer() *Float32Buffer {
	newB := &Float32Buffer{}
	newB.Data = make([]float32, len(buf.Data))
	max := int64(0)
	min := int64(0)
	// try to guess the bit depths without knowing the source
	if buf.SourceBitDepth == 0 {
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
		buf.SourceBitDepth = 8
		if max > 255 || min < 0 {
			buf.SourceBitDepth = 16
		}
		// greater than int16, expecting int24
		if max > 1<<15 {
			buf.SourceBitDepth = 24
		}
		// int 32
		if max > 1<<23 {
			buf.SourceBitDepth = 32
		}
		// int 64
		if max > 1<<31 {
			buf.SourceBitDepth = 64
		}
	}
	newB.SourceBitDepth = buf.SourceBitDepth
	var toFloat func(int) float32
	if buf.SourceBitDepth == 8 {
		// 8-bit uses unsigned ints
		toFloat = func(d int) float32 {
			return float32(d)/255*2 - 1
		}
	} else {
		factor := 1.0 / math.Pow(2, float64(buf.SourceBitDepth-1))
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
