package ffmpeg

import (
	"os/exec"
	"strings"
)

type Command struct {
	Cfg     *Config
	Path    string
	Inputs  []*Video
	Outputs []*Video
	// back    *Background
	// videos  []*Video
	// audio   []*Audio
	// texts   []*Text
}

func NewCommand(path string, opts ...Option) (*Command, error) {
	if strings.TrimSpace(path) == "" {
		path = "ffmpeg"
	}

	cfg := defaultConfig
	for _, opt := range opts {
		opt(cfg)
	}

	cmd := &Command{
		Cfg:  cfg,
		Path: path,
	}

	if err := cmd.Check(); err != nil {
		return nil, err
	}
	return cmd, nil
}

func (f *Command) Check() error {
	return exec.Command(f.Path, "-version").Run()
}
