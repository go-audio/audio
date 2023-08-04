package audio

import (
	"testing"
)

func TestIntBuffer_AsFloat32Buffer(t *testing.T) {
	type fields struct {
		Range          []int
		SourceBitDepth int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "8bit range", // [0, 255]
			fields: fields{Range: []int{0, 1<<8 - 1}, SourceBitDepth: 8}},
		{name: "16bit range", // [-32768, 32767]
			fields: fields{Range: []int{-1<<15, 1<<15 - 1}, SourceBitDepth: 16}},
		{name: "24bit range", // [-8388608, 8388607]
			fields: fields{Range: []int{-1<<23, 1<<23 - 1}, SourceBitDepth: 24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &IntBuffer{
				Format:         FormatMono44100,
				SourceBitDepth: tt.fields.SourceBitDepth,
			}
			intData := []int{
				tt.fields.Range[0],
				0,
				tt.fields.Range[1],
			}
			buf.Data = intData
			got := buf.AsFloat32Buffer()
			for i, f := range got.Data {
				if f < -1.0 || f > 1.0 {
					t.Errorf("%d was converted out of range to %f", intData[i], f)
				}
			}
		})
	}
}

func TestIntBuffer_GetSourceBitDepth(t *testing.T) {
	type fields struct {
		Range          []int
		SourceBitDepth int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "empty buf",
			fields: fields{Range: []int{0, 0}, SourceBitDepth: 16}},
		{name: "signed byte",
			fields: fields{Range: []int{-128, 127}, SourceBitDepth: 16}},
		{name: "8bit range", // [0, 255]
			fields: fields{Range: []int{0, 1<<8 - 1}, SourceBitDepth: 8}},
		{name: "16bit range", // [-32768, 32767]
			fields: fields{Range: []int{-1<<15, 1<<15 - 1}, SourceBitDepth: 16}},
		{name: "24bit range", // [-8388608, 8388607]
			fields: fields{Range: []int{-1<<23, 1<<23 - 1}, SourceBitDepth: 24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &IntBuffer{}
			intData := []int{
				tt.fields.Range[0],
				0,
				tt.fields.Range[1],
			}
			buf.Data = intData
			got := buf.GetSourceBitDepth()
			if got != tt.fields.SourceBitDepth {
				t.Errorf("%d was misestimated as %d", tt.fields.SourceBitDepth, got)
			}
		})
	}
}
