package ase

type TagsChunk struct {
	Tags uint16  `label:"Number of tags"`
	_    [8]byte `label:"For future use"`
}

type Tag struct {
	From                   uint16  `label:"From frame"`
	To                     uint16  `label:"To frame"`
	LoopAnimationDirection byte    `label:"Loop animation direction"`
	Repeat                 uint16  `label:"Repeat N Times"`
	_                      [6]byte `label:"For future use"`
	Color                  [3]byte `label:"RGB values of the tag color"`
	_                      byte    `label:"Extra byte"`
}
