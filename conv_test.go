package audio

import "testing"
import "bytes"

func TestInt24BETo32(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  int32
	}{
		{"max", []byte{0x7F, 0xFF, 0xFF}, 8388607},
		{"mid", []byte{0xFF, 0xFF, 0xFF}, -1},
		{"min", []byte{0x80, 0x00, 0x01}, -8388607},
		{"random", []byte{0x5D, 0xCB, 0xED}, 6147053},
		{"random inverted", []byte{0xA2, 0x34, 0x13}, -6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int24BETo32(tt.bytes); got != tt.want {
				t.Errorf("Int24To32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32toInt24BEBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
		val  int32
	}{
		{name: "mid", want: []byte{0xFF, 0xFF, 0xFF}, val: -1},
		{name: "max", want: []byte{0x7F, 0xFF, 0xFF}, val: 8388607},
		{name: "min", want: []byte{0x80, 0x00, 0x01}, val: -8388607},
		{name: "random", want: []byte{0x5D, 0xCB, 0xED}, val: 6147053},
		{name: "random inverted", want: []byte{0xA2, 0x34, 0x13}, val: -6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32toInt24BEBytes(tt.val); bytes.Compare(tt.want, got) != 0 {
				t.Errorf("Int32toInt24BEBytes(%d) = %x, want %x", tt.val, got, tt.want)
			}
		})
	}
}
