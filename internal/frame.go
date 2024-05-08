package internal

type Frame struct {
	Duration uint16
	Layers   []Layer
}

type PreProcessedFrame struct {
	Duration uint16
	Chunks   []any
}
