package ffmpeg

type Video struct {
	Device string
	Format string
	Scala  *Size
	Rate   uint32
	Position
	Size
}
