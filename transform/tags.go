package transform

import (
	"image/color"
	"io"

	"github.com/SnareChops/aseprite-loader/ase"
	"github.com/SnareChops/aseprite-loader/internal"
	"github.com/SnareChops/aseprite-loader/trace"
)

func transformTags(in io.Reader, out io.Writer) (data []internal.Tag, err error) {
	trace.Log("transformTags")
	var tagsHeader ase.TagsChunk
	tagsHeader, err = transform[ase.TagsChunk](in, out)
	if err != nil {
		return
	}
	for range tagsHeader.Tags {
		var tag ase.Tag
		tag, err = transform[ase.Tag](in, out)
		if err != nil {
			return
		}
		var name ase.String
		name, err = transformString(in, out)
		if err != nil {
			return
		}
		data = append(data, internal.Tag{
			Name:          name.Value,
			From:          tag.From,
			To:            tag.To,
			LoopDirection: tag.LoopAnimationDirection,
			Repeat:        tag.Repeat,
			Color:         color.RGBA{tag.Color[0], tag.Color[1], tag.Color[2], 0xff},
		})
	}
	return
}
