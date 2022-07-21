package ffmpeg

type CodecType string

const (
	Audio    CodecType = "audio"
	Video    CodecType = "video"
	Subtitle CodecType = "subtitle"
)
