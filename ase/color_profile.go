package ase

type ColorProfileChunk struct {
	Type  uint16  `label:"Type"`
	Flags uint16  `label:"Flags"`
	Gamma [4]byte `label:"Fixed Gamma"`
	_     [8]byte `label:"Reserved"`
}

type ColorProfileICC struct {
	Length uint32 `label:"ICC Profile Data Length"`
	Data   []byte `label:"ICC Profile Data"`
}
