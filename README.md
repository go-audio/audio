# audio

`audio` is a generic Go package designed to define a common
interface to analyze and/or process audio data.

At the heart of the package is the `Buffer` interface and its implementations:

* `FloatBuffer`
* `IntBuffer`

Decoders, encoders, processors, analyzers and transformers can be written to
accept or return these types and share a common interface.