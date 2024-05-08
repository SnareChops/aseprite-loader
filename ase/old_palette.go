package ase

type OldPaletteHeader struct {
	NumberOfPackets uint16 `label:"Number of Packets"`
}

type OldPalettePacket struct {
	Skip           byte `label:"Number of palette entries to skip from the last packet (start from 0)"`
	NumberOfColors byte `label:"Number of colors in the packet (0 means 256)"`
}

type OldPacketColor struct {
	Red   byte `label:"Red"`
	Green byte `label:"Green"`
	Blue  byte `label:"Blue"`
}
