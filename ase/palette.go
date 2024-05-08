package ase

type PaletteChunk struct {
	Size       uint32  `label:"New palette size (total number of entries)"`
	FirstIndex uint32  `label:"Fist color index to change"`
	LastIndex  uint32  `label:"Last color index to change"`
	_          [8]byte `label:"For future use"`
}

type PaletteEntry struct {
	Flags uint16 `label:"Entry flags"`
	Red   byte   `label:"Red"`
	Green byte   `label:"Green"`
	Blue  byte   `label:"Blue"`
	Alpha byte   `label:"Alpha"`
}
