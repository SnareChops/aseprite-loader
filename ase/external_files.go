package ase

type ExternalFilesChunk struct {
	Entries uint32  `label:"Number of entries"`
	_       [8]byte `label:"Reserved"`
}

type ExternalFileEntry struct {
	EntryID uint32  `label:"Entry ID"`
	Type    byte    `label:"Type"`
	_       [7]byte `label:"Reserved"`
}
