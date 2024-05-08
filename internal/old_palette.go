package internal

import "image/color"

type OldPalette struct {
	Packets  []OldPacket
	UserData UserData
}

type OldPacket struct {
	Colors []color.Color
}
