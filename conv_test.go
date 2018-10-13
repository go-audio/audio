package audio

import (
	"bytes"
	"math"
	"testing"
)

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
				t.Errorf("Int24BETo32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt24LETo32(t *testing.T) {
	tests := []struct {
		name  string
		bytes []byte
		want  int32
	}{
		{"max", []byte{0xFF, 0xFF, 0x7F}, 8388607},
		{"mid", []byte{0xFF, 0xFF, 0xFF}, -1},
		{"min", []byte{0x01, 0x00, 0x80}, -8388607},
		{"random", []byte{0xED, 0xCB, 0x5D}, 6147053},
		{"random inverted", []byte{0x13, 0x34, 0xA2}, -6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int24LETo32(tt.bytes); got != tt.want {
				t.Errorf("Int24LETo32() = %v, want %v", got, tt.want)
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

func TestInt32toInt24LEBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
		val  int32
	}{
		{name: "mid", want: []byte{0xFF, 0xFF, 0xFF}, val: -1},
		{name: "max", want: []byte{0xFF, 0xFF, 0x7F}, val: 8388607},
		{name: "min", want: []byte{0x01, 0x00, 0x80}, val: -8388607},
		{name: "random", want: []byte{0xED, 0xCB, 0x5D}, val: 6147053},
		{name: "random inverted", want: []byte{0x13, 0x34, 0xA2}, val: -6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32toInt24LEBytes(tt.val); bytes.Compare(tt.want, got) != 0 {
				t.Errorf("Int32toInt24LEBytes(%d) = %x, want %x", tt.val, got, tt.want)
			}
		})
	}
}

// round tripping similar to aiff to wav
func TestInt24BETo32ToLEToInt32(t *testing.T) {
	tests := []struct {
		name string
		be   []byte
		le   []byte
		val  int32
	}{
		{name: "mid", be: []byte{0xFF, 0xFF, 0xFF}, le: []byte{0xff, 0xff, 0xff}, val: -1},
		{name: "max", be: []byte{0x7F, 0xFF, 0xFF}, le: []byte{0xff, 0xff, 0x7f}, val: 8388607},
		{name: "min", be: []byte{0x80, 0x00, 0x01}, le: []byte{0x01, 0x00, 0x80}, val: -8388607},
		{name: "random", be: []byte{0x5D, 0xCB, 0xED}, le: []byte{0xED, 0xCB, 0x5D}, val: 6147053},
		{name: "random inverted", be: []byte{0xA2, 0x34, 0x13}, le: []byte{0x13, 0x34, 0xA2}, val: -6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beV := Int24BETo32(tt.be)
			if beV != tt.val {
				t.Errorf("Int24BETo32(%x) = %d, want %d", tt.be, beV, tt.val)
			}
			leB := Int32toInt24LEBytes(beV)
			if bytes.Compare(leB, tt.le) != 0 {
				t.Errorf("Int32toInt24LEBytes(%d) = %#v, want %#v", beV, leB, tt.le)
			}
			leV := Int24LETo32(leB)
			if leV != tt.val {
				t.Errorf("Int24LETo32(%#v) = %d, want %d", leB, leV, tt.val)
			}
		})
	}
}

func TestUint32toUint24Bytes(t *testing.T) {
	tests := []struct {
		name string
		be   []byte
		val  uint32
	}{
		{name: "mid", be: []byte{0x0, 0x0, 0x1}, val: 1},
		{name: "max", be: []byte{0xFF, 0xFF, 0xFF}, val: math.MaxUint32},
		{name: "random", be: []byte{0x5D, 0xCB, 0xED}, val: 6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beB := Uint32toUint24Bytes(tt.val)
			if bytes.Compare(beB, tt.be) != 0 {
				t.Errorf("Uint32toUint24Bytes(%d) = %x, want %x", tt.val, beB, tt.be)
			}
		})
	}
}

func TestUint24to32(t *testing.T) {
	tests := []struct {
		name string
		be   []byte
		val  uint32
	}{
		{name: "mid", be: []byte{0x0, 0x0, 0x1}, val: 1},
		{name: "max", be: []byte{0xff, 0xff, 0xff}, val: 16777215},
		{name: "random", be: []byte{0x5D, 0xCB, 0xED}, val: 6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := Uint24to32(tt.be)
			if val != tt.val {
				t.Errorf("Uint24to32(%x) = %d, want %d", tt.be, val, tt.val)
			}
		})
	}
}

func TestIntToIEEEFloat(t *testing.T) {
	tests := []struct {
		name string
		ret  [10]byte
		val  int
		val2 int
	}{
		{name: "min", ret: [10]byte{0x3f, 0xff, 0x80}, val: 1, val2: 1},
		{name: "max", ret: [10]byte{0x40, 0x3e, 0x80}, val: math.MaxInt64, val2: 800000000}, // IEEEFloatToInt truncates
		{name: "random", ret: [10]byte{0x40, 0x15, 0xbb, 0x97, 0xda}, val: 6147053, val2: 6147053},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := IntToIEEEFloat(tt.val)
			if val != tt.ret {
				t.Errorf("IntToIEEEFloat(%d) = %x, want %x", tt.val, val, tt.ret)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := IEEEFloatToInt(tt.ret)
			if val != tt.val2 {
				t.Errorf("IEEEFloatToInt(%x) = %d, want %d", tt.ret, val, tt.val2)
			}
		})
	}
}
