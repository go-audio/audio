package audio

import (
	"reflect"
	"testing"
)

func TestPCMBuffer_AsFloatBuffer(t *testing.T) {
	tests := []struct {
		name     string
		datatype PCMDataFormat
		bd       uint8
		i8       []int8
		i16      []int16
		i32      []int32
		f32      []float32
		f64      []float64
	}{
		{"int8 conversion", DataTypeI8, 1, []int8{1, 2, 3}, []int16{1, 2, 3}, []int32{1, 2, 3}, []float32{2, 4, 6}, []float64{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pcmb := PCMBuffer{
				Format:         FormatMono22500,
				I8:             tt.i8,
				SourceBitDepth: tt.bd,
				DataType:       tt.datatype,
			}

			// I'm unsure if the actual behavior is for int8 individual
			// data to double the way it did
			// 1,2,3  => 2,4,6
			pcmb32 := pcmb.AsFloat32Buffer()
			if !reflect.DeepEqual(pcmb32.Data, tt.f32) {
				t.Errorf("Expected %+v got %+v", tt.f32, pcmb32.Data)
			}
			pcmb64 := pcmb.AsFloatBuffer()
			if !reflect.DeepEqual(pcmb64.Data, tt.f64) {
				t.Errorf("Expected %+v got %+v", tt.f64, pcmb64.Data)
			}

			if !reflect.DeepEqual(pcmb.AsI16(), tt.i16) {
				t.Errorf("Expected %+v got %+v", tt.i8, pcmb.AsI16())
			}
			if !reflect.DeepEqual(pcmb.AsI32(), tt.i32) {
				t.Errorf("Expected %+v got %+v", tt.i32, pcmb.AsI32())
			}
		})
	}

}
