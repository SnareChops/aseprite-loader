package ase

type CelExtraChunk struct {
	Flags     uint32   `label:"Flags"`
	XPosition [4]byte  `label:"Precise X position"`
	YPosition [4]byte  `label:"Precise Y position"`
	Width     [4]byte  `label:"Width of the cel in the sprite"`
	Height    [4]byte  `label:"Height of the cel in the sprite"`
	_         [16]byte `label:"For future use"`
}
