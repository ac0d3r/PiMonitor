package raspivid

import (
	"context"
	"os/exec"
	"strconv"
)

var (
	BinPath string = "raspivid"
)

type Options struct {
	FPS            int
	Width          int
	Height         int
	Rotation       int
	HorizontalFlip bool
	VerticalFlip   bool
}

func (opt Options) MakeArgs() []string {
	args := []string{
		"-ih",
		"-t", "0", // no timeout
		"-fps", strconv.Itoa(opt.FPS), // frames/sec
		"-w", strconv.Itoa(opt.Width), // width
		"-h", strconv.Itoa(opt.Height), // height
		"-pf", "baseline",
		"-n",
		"-o", "-",
	}
	if opt.HorizontalFlip {
		args = append(args, "--hflip")
	}
	if opt.VerticalFlip {
		args = append(args, "--vflip")
	}
	if opt.Rotation != 0 {
		args = append(args, "--rotation", strconv.Itoa(opt.Rotation))
	}

	return args
}

func makeCmd(ctx context.Context, opt Options) *exec.Cmd {
	return exec.CommandContext(ctx, BinPath, opt.MakeArgs()...)
}
