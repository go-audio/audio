package audio

import (
	"math"
	"reflect"
	"testing"
)

func TestFloat64Buffer(t *testing.T) {
	tests := []struct {
		name    string
		f64     []float64
		f32     []float32
		integer []int
	}{
		{"float 64 conversion", []float64{1, 2, 3}, []float32{1, 2, 3}, []int{1, 2, 3}},
		{"float 64 conversion, max float32", []float64{1, 2, float64(math.MaxFloat32)}, []float32{1, 2, math.MaxFloat32}, []int{1, 2, int(math.Inf(1))}},
		{"float 64 conversion, inf", []float64{1, 2, math.MaxFloat64}, []float32{1, 2, float32(math.Inf(1))}, []int{1, 2, int(math.Inf(1))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := FloatBuffer{Format: FormatMono22500, Data: tt.f64}
			fb32 := fb.AsFloat32Buffer()
			if !reflect.DeepEqual(fb32.Data, tt.f32) {
				t.Errorf("Expected %+v got %+v", tt.f32, fb32.Data)
			}
			integer := fb.AsIntBuffer()
			if !reflect.DeepEqual(integer.Data, tt.integer) {
				t.Errorf("Expected %+v got %+v", tt.integer, integer.Data)
			}
		})
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		name    string
		f64     []float64
		f32     []float32
		integer []int
	}{
		{"float 64 conversion", []float64{1, 2, 3}, []float32{1, 2, 3}, []int{1, 2, 3}},
		{"float 64 conversion, max float32", []float64{1, 2, float64(math.MaxFloat32)}, []float32{1, 2, math.MaxFloat32}, []int{1, 2, int(math.Inf(1))}},
		{"float 64 conversion, inf", []float64{1, 2, math.MaxFloat64}, []float32{1, 2, float32(math.Inf(1))}, []int{1, 2, int(math.Inf(1))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := FloatBuffer{Format: FormatMono22500, Data: tt.f64}
			b := fb.Clone()
			fb2 := b.AsFloatBuffer()
			if !reflect.DeepEqual(fb2.Data, tt.f64) {
				t.Errorf("Expected %+v got %+v", tt.f64, fb2.Data)
			}
			fb3 := b.AsFloat32Buffer()
			b = fb3.Clone()
			if !reflect.DeepEqual(fb3.Data, tt.f32) {
				t.Errorf("Expected %+v got %+v", tt.f32, fb3.Data)
			}
			fb4 := b.AsIntBuffer()
			b = fb4.Clone()
			if !reflect.DeepEqual(fb4.Data, tt.integer) {
				t.Errorf("Expected %+v got %+v", tt.integer, fb4.Data)
			}
		})
	}
}

func TestNumFrames(t *testing.T) {
	expect := 3
	fb64 := FloatBuffer{Format: FormatMono22500, Data: []float64{1, 2, 3}}
	numFrames := fb64.NumFrames()
	if numFrames != expect {
		t.Errorf("Expected %d got %d", expect, numFrames)
	}
	fb32 := Float32Buffer{Format: FormatMono22500, Data: []float32{1, 2, 3}}
	numFrames = fb32.NumFrames()
	if numFrames != expect {
		t.Errorf("Expected %d got %d", expect, numFrames)
	}
}

func TestFloat32Buffer(t *testing.T) {
	tests := []struct {
		name    string
		f64     []float64
		f32     []float32
		integer []int
	}{
		{"float 32 conversion", []float64{1, 2, 3}, []float32{1, 2, 3}, []int{1, 2, 3}},
		{"float 32 conversion, max float32", []float64{1, 2, float64(math.MaxFloat32)}, []float32{1, 2, math.MaxFloat32}, []int{1, 2, int(math.Inf(1))}},
		{"float 32 conversion, inf", []float64{1, 2, math.Inf(1)}, []float32{1, 2, float32(math.Inf(1))}, []int{1, 2, int(math.Inf(1))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := Float32Buffer{Format: FormatMono22500, Data: tt.f32}
			fb64 := fb.AsFloatBuffer()
			if !reflect.DeepEqual(fb64.Data, tt.f64) {
				t.Errorf("Expected %+v got %+v", tt.f64, fb64.Data)
			}
			integer := fb.AsIntBuffer()
			if !reflect.DeepEqual(integer.Data, tt.integer) {
				t.Errorf("Expected %+v got %+v", tt.integer, integer.Data)
			}
		})
	}
}
