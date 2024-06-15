package internal

type Frame struct {
	Duration uint16
	Cels     []Cel
}

type PreProcessedFrame struct {
	Duration uint16
	Chunks   []any
}
