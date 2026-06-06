package video

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os/exec"
)

type Decoder struct {
	cmd *exec.Cmd
	stdout io.ReadCloser
	reader *bufio.Reader
}

func NewDecoder(videoPath string, width, height, fps int) (*Decoder) {
	filer := fmt.Sprintf("fps=%d, scale=%d:%d", fps, width, height)

	cmd := exec.Command(
		"ffmpeg", 
		"-i", videoPath, 
		"-vf", filter, 
		"-f", "image2pipe", 
		"-vcodec","mjpeg", 
		"-loglevel","quiet", 
		"pipe:1")
	return &Decoder{
		cmd: cmd,
	}
}

func (d *Decoder) Start() error {
	var err error
	d.stdout, err = d.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("stdout pipe got cooked: %w", err)
	}

	if err := d.cmd.Start(); err != nil {
		return fmt.Errorf("ffmpeg got cooked: %w", err)
	}

	d.reader = bufio.NewReader(d.stdout)
	return nil
}

func (d *Decoder) NextFrame() (image.Image, error) {
	img, _, err := image.Decode(d.reader)
	if err != nil {
		return nil, fmt.Errorf("decoder got cooked gang: %w", err)
	}
	return img, nil
}

func (d *Decoder) Stop() error {
	if d.stdout != nil {
		d.stdout.Close()
	}

	if d.cmd != nil && d.cmd.Process != nil{
		return d.cmd.Process.Kill()
	}

	return nil
}

