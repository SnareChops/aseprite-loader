package internal

type ColorProfile struct {
	Type     uint16
	Flags    uint16
	Gamma    [4]byte
	ICC      []byte
	UserData UserData
}
