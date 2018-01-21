package audio

import (
	"testing"
)

func TestIntBuffer_AsFloat32Buffer(t *testing.T) {
	type fields struct {
		Range          int
		SourceBitDepth int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "16bit range",
			fields: fields{Range: int(int16(1<<15 - 1)), SourceBitDepth: 16}},
		{name: "24bit range",
			fields: fields{Range: int(int32(1 << 23)), SourceBitDepth: 24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &IntBuffer{
				Format:         FormatMono44100,
				SourceBitDepth: tt.fields.SourceBitDepth,
			}
			intData := []int{
				-tt.fields.Range,
				0,
				tt.fields.Range,
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
