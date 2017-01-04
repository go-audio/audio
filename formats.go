package audio

var (
	// MONO

	// FormatMono22500 mono 22.5kHz format.
	FormatMono22500 = &Format{
		NumChannels: 1,
		SampleRate:  22500,
	}
	// FormatMono44100 mono 8bit 44.1kHz format.
	FormatMono44100 = &Format{
		NumChannels: 1,
		SampleRate:  44100,
	}
	// FormatMono48000 mono 48kHz format.
	FormatMono48000 = &Format{
		NumChannels: 1,
		SampleRate:  48000,
	}
	// FormatMono96000 mono 96kHz format.
	FormatMono96000 = &Format{
		NumChannels: 1,
		SampleRate:  96000,
	}

	// STEREO

	// FormatStereo22500 Stereo 22.5kHz format.
	FormatStereo22500 = &Format{
		NumChannels: 1,
		SampleRate:  22500,
	}
	// FormatStereo44100 Stereo 8bit 44.1kHz format.
	FormatStereo44100 = &Format{
		NumChannels: 1,
		SampleRate:  44100,
	}
	// FormatStereo48000 Stereo 48kHz format.
	FormatStereo48000 = &Format{
		NumChannels: 1,
		SampleRate:  48000,
	}
	// FormatStereo96000 Stereo 96kHz format.
	FormatStereo96000 = &Format{
		NumChannels: 1,
		SampleRate:  96000,
	}
)
