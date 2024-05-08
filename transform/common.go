package transform

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformString(in io.Reader, out io.Writer) (data ase.String, err error) {
	trace.Log("transformString")
	err = binary.Read(in, binary.LittleEndian, &data.Length)
	if err != nil {
		return
	}
	if data.Length > 0 {
		value := make([]byte, data.Length)
		in.Read(value)
		data.Value = string(value)
	}
	err = write(out, data)
	return
}

func transformBytes[T int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint](in io.Reader, out io.Writer, length T) (data []byte, err error) {
	trace.Log("transformBytes")
	data = make([]byte, length)
	_, err = in.Read(data)
	if err != nil {
		return
	}
	err = write(out, data)
	return
}

func decompress(compressed []byte) (data []byte, err error) {
	trace.Log("decompress")
	reader := bytes.NewReader(compressed)
	var zr io.ReadCloser
	zr, err = zlib.NewReader(reader)
	if err != nil {
		return
	}
	defer zr.Close()
	data, err = io.ReadAll(zr)
	return
}
