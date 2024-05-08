package ase

type FrameHeader struct {
	FrameSize     uint32  `label:"Frame Size"`
	Magic         uint16  `label:"Magic Number"`
	OldNumChunks  uint16  `label:"Old Number of Chunks"`
	FrameDuration uint16  `label:"Frame Duration"`
	_             [2]byte `label:"For Future Use"`
	NewNumChunks  uint32  `label:"New Number of Chunks"`
}
