package ffmpeg

import "os/exec"

type Ffmpeg struct {
	Path string
}

func (f *Ffmpeg) Check(path string) error {
	return exec.Command(path, "-version").Run()
}
