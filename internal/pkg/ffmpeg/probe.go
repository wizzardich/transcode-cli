package ffmpeg

import (
	"encoding/json"
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type ProbeTags struct {
	Language string `json:"language"`
	Title    string `json:"title"`
}

type ProbeDisposition struct {
	Default  int `json:"default"`
	Dub      int `json:"dub"`
	Original int `json:"original"`
	Comment  int `json:"comment"`
	Forced   int `json:"forced"`
}

type ProbeStream struct {
	Index         int              `json:"index"`
	Channels      int              `json:"channels"`
	ChannelLayout string           `json:"channel_layout"`
	CodecName     string           `json:"codec_name"`
	CodecType     CodecType        `json:"codec_type"`
	BitRate       string           `json:"bit_rate"`
	Tags          ProbeTags        `json:"tags"`
	Disposition   ProbeDisposition `json:"disposition"`
}

type ProbeFormat struct {
	Filename     string `json:"filename"`
	StreamNumber int    `json:"nb_streams"`
	FormatName   string `json:"format_name"`
	Duration     string `json:"duration"`
}

type Probe struct {
	Streams []ProbeStream `json:"streams"`
	Format  ProbeFormat   `json:"format"`
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}

func (t *Target) Describe() string {
	result := fmt.Sprintf("%s streams:", t.probe().Format.Filename)

	videoStreams := t.VideoStreams()
	if len(videoStreams) > 0 {
		result += "\n    Video streams:"
		for _, stream := range videoStreams {
			result += "\n        - " + stream.Describe()
		}
	}

	audioStreams := t.AudioStreams()
	if len(audioStreams) > 0 {
		result += "\n    Audio streams:"
		for _, stream := range audioStreams {
			result += "\n        - " + stream.Describe()
		}
	}

	subtitleStreams := t.SubtitleStreams()
	if len(subtitleStreams) > 0 {
		result += "\n    Subtitle streams:"
		for _, stream := range subtitleStreams {
			result += "\n        - " + stream.Describe()
		}
	}

	return result
}

func (t *Target) probe() Probe { // TODO: add cache
	probeJson, err := ffmpeg.Probe(t.Path)
	fail(err)
	probe := Probe{}
	err = json.Unmarshal([]byte(probeJson), &probe)
	fail(err)
	return probe
}

func (t *Target) streams(filter func(ProbeStream) bool) []ProbeStream {
	probe := t.probe()

	filteredStreams := []ProbeStream{}
	for _, stream := range probe.Streams {
		if filter(stream) {
			filteredStreams = append(filteredStreams, stream)
		}
	}

	return filteredStreams
}

func (t *Target) AudioStreams() []ProbeStream {
	return t.streams(func(stream ProbeStream) bool { return stream.CodecType == Audio })
}

func (t *Target) VideoStreams() []ProbeStream {
	return t.streams(func(stream ProbeStream) bool { return stream.CodecType == Video })
}

func (t *Target) SubtitleStreams() []ProbeStream {
	return t.streams(func(stream ProbeStream) bool { return stream.CodecType == Subtitle })
}

func (p *ProbeStream) Describe() string {

	result := fmt.Sprintf("#%d -- %s", p.Index, p.CodecType)
	if p.Tags.Language != "" {
		result += fmt.Sprintf("(%s)", p.Tags.Language)
	}
	if p.Tags.Title != "" {
		result += fmt.Sprintf(" \"%s\"", p.Tags.Title)
	}

	if p.CodecType == Video {
		result += fmt.Sprintf(" stream (%s)", p.CodecName)
	}

	if p.CodecType == Audio {
		result += fmt.Sprintf(" stream (%d channels, %s, %s)", p.Channels, p.CodecName, p.BitRate)
	}

	if p.CodecType == Subtitle {
		result += " stream"
		if p.Disposition.Forced == 1 {
			result += " (forced)"
		}
		if p.Disposition.Comment == 1 {
			result += " (comment)"
		}
	}
	return result
}
