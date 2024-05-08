package ase

type LayerChunk struct {
	Flags         uint16  `label:"Flags"`
	Type          uint16  `label:"Layer Type"`
	ChildLevel    uint16  `label:"Child Level"`
	DefaultWidth  uint16  `label:"Default Layer Width"`
	DefaultHeight uint16  `label:"Default Layer Height"`
	BlendMode     uint16  `label:"Blend Mode"`
	Opacity       byte    `label:"Opacity"`
	_             [3]byte `label:"For Future Use"`
}
