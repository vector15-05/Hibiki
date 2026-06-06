package audio

import (
	"os/exec"
)

type Player struct {
	cmd *exec.Cmd
}

func NewPlayer(audioPath string) *Player{
	cmd := exec.Command("mpv", "--no-video", "--really-quiet", audioPath)
	return &Player{cmd: cmd}
}

func (p *Player) Start() error {
	return p.cmd.Start()
}

func (p *Player) Stop() error{
	if p.cmd != nil && p.cmd.Process != nil {
		return p.cmd.Process.Kill()
	}
	return nil
}