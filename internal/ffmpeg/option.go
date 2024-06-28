package ffmpeg

import "runtime"

// 调优
type (
	Config struct {
		VideoEncoder string
		AudioEncoder string
		Preset       string
		Tune         string
		Threads      int
		BufSize      int // uints: M
		RTBufSize    int // uints: M
	}
	Option func(cfg *Config)
)

var defaultConfig = &Config{
	VideoEncoder: "libx264", // libx264 | libx265 | libvpx | libvpx-vp9 | libaom-av1 | mpeg2video | mpeg4
	AudioEncoder: "acc",     // aac | libmp3lame | libopus | libvorbis | libflac | alac
	Preset:       "medium",  // ultrafast | superfast | veryfast | faster | fast | medium | slow | slower | veryslow | placebo
	Tune:         "",        // film | animation | grain | stillimage | fastdecode | zerolatency | psnr | ssim
	Threads:      runtime.NumCPU() / 2,
	BufSize:      2,
	RTBufSize:    10,
}

func WithVideoEncoder(encoder string) Option {
	return func(cfg *Config) {
		cfg.VideoEncoder = encoder
	}
}

func WithAudioEncoder(encoder string) Option {
	return func(cfg *Config) {
		cfg.AudioEncoder = encoder
	}
}

func WithPreset(preset string) Option {
	return func(cfg *Config) {
		cfg.Preset = preset
	}
}

func WithTune(tune string) Option {
	return func(cfg *Config) {
		cfg.Tune = tune
	}
}

func WithThreads(threads int) Option {
	return func(cfg *Config) {
		cfg.Threads = threads
	}
}

func WithBufSize(bufSize int) Option {
	return func(cfg *Config) {
		cfg.BufSize = bufSize
	}
}

func WithRTBufSize(bufSize int) Option {
	return func(cfg *Config) {
		cfg.RTBufSize = bufSize
	}
}
