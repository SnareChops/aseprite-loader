package internal

type CelExtra struct {
	Flags    uint32
	X        [4]byte
	Y        [4]byte
	Width    [4]byte
	Height   [4]byte
	UserData UserData
}
