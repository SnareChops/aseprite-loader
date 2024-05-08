package transform

import (
	"io"

	"github.com/SnareChops/aseprite/ase"
	"github.com/SnareChops/aseprite/internal"
	"github.com/SnareChops/aseprite/trace"
)

func transformExternalFiles(in io.Reader, out io.Writer) (data internal.ExternalFiles, err error) {
	trace.Log("transformExternalFiles")
	var externalFiles ase.ExternalFilesChunk
	externalFiles, err = transform[ase.ExternalFilesChunk](in, out)
	if err != nil {
		return
	}
	for range externalFiles.Entries {
		var entry ase.ExternalFileEntry
		entry, err = transform[ase.ExternalFileEntry](in, out)
		if err != nil {
			return
		}
		var name ase.String
		name, err = transformString(in, out)
		if err != nil {
			return
		}
		data.Entries = append(data.Entries, internal.ExternalFile{
			EntryID: entry.EntryID,
			Type:    entry.Type,
			Name:    name.Value,
		})
	}
	return
}
