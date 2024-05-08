package internal

type Mask struct {
	X        int16
	Y        int16
	Width    uint16
	Height   uint16
	Name     string
	Data     []byte
	UserData UserData
}
