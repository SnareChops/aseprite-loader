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
	trace.Log("transformBytes", length)
	data = make([]byte, length)
	var total T = 0
	for total < length {
		var n int
		n, err = in.Read(data[total:])
		if err != nil {
			return
		}
		total += T(n)
	}
	err = write(out, data)
	return
}

func decompress(compressed []byte) (data []byte, err error) {
	trace.Log("decompress", len(compressed))
	reader := bytes.NewReader(compressed)
	trace.Log("reader length", reader.Len())
	var zr io.ReadCloser
	zr, err = zlib.NewReader(reader)
	if err != nil {
		panic(err)
		return
	}
	defer zr.Close()
	data, err = io.ReadAll(zr)
	if err != nil {
		panic(err)
	}
	return
}
