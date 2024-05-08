package ase

type MaskChunk struct {
	X      int16   `label:"X position"`
	Y      int16   `label:"Y position"`
	Width  uint16  `label:"Width"`
	Height uint16  `label:"Height"`
	_      [8]byte `label:"For future use"`
}
