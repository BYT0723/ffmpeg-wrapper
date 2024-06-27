package ffmpeg

import "os/exec"

type Ffmpeg struct {
	Path   string
	back   *Background
	videos []*Video
	audio  []*Audio
	texts  []*Text
}

func (f *Ffmpeg) Check(path string) error {
	return exec.Command(path, "-version").Run()
}
