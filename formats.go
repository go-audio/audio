package audio

import "encoding/binary"

var (
	// AIFF
	// MONO

	// FormatMono225008bBE mono 8bit 22.5kHz AIFF like format.
	FormatMono225008bBE = &Format{
		NumChannels: 1,
		SampleRate:  22500,
		BitDepth:    8,
		Endianness:  binary.BigEndian,
	}
	// FormatMono2250016bBE mono 16bit 22.5kHz AIFF like format.
	FormatMono2250016bBE = &Format{
		NumChannels: 1,
		SampleRate:  22500,
		BitDepth:    16,
		Endianness:  binary.BigEndian,
	}
	// FormatMono441008bBE mono 8bit 44.1kHz AIFF like format.
	FormatMono441008bBE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    8,
		Endianness:  binary.BigEndian,
	}
	// FormatMono4410016bBE mono 16bit 44.1kHz AIFF like format.
	FormatMono4410016bBE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    16,
		Endianness:  binary.BigEndian,
	}
	// FormatMono4410024bBE mono 24bit 44.1kHz AIFF like format.
	FormatMono4410024bBE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    24,
		Endianness:  binary.BigEndian,
	}
	// FormatMono4410032bBE mono 32bit 44.1kHz AIFF like format.
	FormatMono4410032bBE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    32,
		Endianness:  binary.BigEndian,
	}

	// STEREO

	// FormatStereo225008bBE Stereo 8bit 22.5kHz AIFF like format.
	FormatStereo225008bBE = &Format{
		NumChannels: 2,
		SampleRate:  22500,
		BitDepth:    8,
		Endianness:  binary.BigEndian,
	}
	// FormatStereo2250016bBE Stereo 16bit 22.5kHz AIFF like format.
	FormatStereo2250016bBE = &Format{
		NumChannels: 2,
		SampleRate:  22500,
		BitDepth:    16,
		Endianness:  binary.BigEndian,
	}
	// FormatStereo441008bBE Stereo 8bit 44.1kHz AIFF like format.
	FormatStereo441008bBE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    8,
		Endianness:  binary.BigEndian,
	}
	// FormatStereo4410016bBE Stereo 16bit 44.1kHz AIFF like format.
	FormatStereo4410016bBE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    16,
		Endianness:  binary.BigEndian,
	}
	// FormatStereo4410024bBE Stereo 24bit 44.1kHz AIFF like format.
	FormatStereo4410024bBE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    24,
		Endianness:  binary.BigEndian,
	}
	// FormatStereo4410032bBE Stereo 32bit 44.1kHz AIFF like format.
	FormatStereo4410032bBE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    32,
		Endianness:  binary.BigEndian,
	}

	// WAV
	// MONO

	// FormatMono225008bLE mono 8bit 22.5kHz AIFF like format.
	FormatMono225008bLE = &Format{
		NumChannels: 1,
		SampleRate:  22500,
		BitDepth:    8,
		Endianness:  binary.LittleEndian,
	}
	// FormatMono2250016bLE mono 16bit 22.5kHz AIFF like format.
	FormatMono2250016bLE = &Format{
		NumChannels: 1,
		SampleRate:  22500,
		BitDepth:    16,
		Endianness:  binary.LittleEndian,
	}
	// FormatMono441008bLE mono 8bit 44.1kHz AIFF like format.
	FormatMono441008bLE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    8,
		Endianness:  binary.LittleEndian,
	}
	// FormatMono4410016bLE mono 16bit 44.1kHz AIFF like format.
	FormatMono4410016bLE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    16,
		Endianness:  binary.LittleEndian,
	}
	// FormatMono4410024bLE mono 24bit 44.1kHz AIFF like format.
	FormatMono4410024bLE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    24,
		Endianness:  binary.LittleEndian,
	}
	// FormatMono4410032bLE mono 32bit 44.1kHz AIFF like format.
	FormatMono4410032bLE = &Format{
		NumChannels: 1,
		SampleRate:  44100,
		BitDepth:    32,
		Endianness:  binary.LittleEndian,
	}

	// STEREO

	// FormatStereo225008bLE Stereo 8bit 22.5kHz AIFF like format.
	FormatStereo225008bLE = &Format{
		NumChannels: 2,
		SampleRate:  22500,
		BitDepth:    8,
		Endianness:  binary.LittleEndian,
	}
	// FormatStereo2250016bLE Stereo 16bit 22.5kHz AIFF like format.
	FormatStereo2250016bLE = &Format{
		NumChannels: 2,
		SampleRate:  22500,
		BitDepth:    16,
		Endianness:  binary.LittleEndian,
	}
	// FormatStereo441008bLE Stereo 8bit 44.1kHz AIFF like format.
	FormatStereo441008bLE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    8,
		Endianness:  binary.LittleEndian,
	}
	// FormatStereo4410016bLE Stereo 16bit 44.1kHz AIFF like format.
	FormatStereo4410016bLE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    16,
		Endianness:  binary.LittleEndian,
	}
	// FormatStereo4410024bLE Stereo 24bit 44.1kHz AIFF like format.
	FormatStereo4410024bLE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    24,
		Endianness:  binary.LittleEndian,
	}
	// FormatStereo4410032bLE Stereo 32bit 44.1kHz AIFF like format.
	FormatStereo4410032bLE = &Format{
		NumChannels: 2,
		SampleRate:  44100,
		BitDepth:    32,
		Endianness:  binary.LittleEndian,
	}
)
