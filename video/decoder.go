package video

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Decoder struct {
	cmd *exec.Cmd
	stdout io.ReadCloser
	width int
	height int
}

func NewDecoder(videoPath string, width, height, fps int) (*Decoder) {
	filter := fmt.Sprintf("fps=%d, scale=%d:%d", fps, width, height)

	cmd := exec.Command(
		"ffmpeg", 
		"-i", videoPath, 
		"-vf", filter, 
		"-f", "rawvideo",     
		"-pix_fmt", "gray",   
		"-loglevel", "quiet", 
		"pipe:1")
	cmd.Stderr = os.Stderr 
	return &Decoder{
		cmd: cmd,
		width:  width,
		height: height,
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

	return nil
}

func (d *Decoder) NextFrame() ([]byte, error) {
	frameSize := d.width * d.height
	buffer := make([]byte, frameSize)
	_, err := io.ReadFull(d.stdout, buffer) // grab the whole frame no more no less

	if err != nil {
		return nil, err
	}
	return buffer, nil
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

