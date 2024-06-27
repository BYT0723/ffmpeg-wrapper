package ffmpeg

type Video struct {
	Position
	Size
	Device string
	Scala  *Size
	Rate   string
}
