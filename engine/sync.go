package engine

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vector15-05/Hibiki/audio"
	"github.com/vector15-05/Hibiki/render"
	"github.com/vector15-05/Hibiki/video"
)

func Run(mediaPath string) error{

	vDecoder := video.NewDecoder(mediaPath, 100, 40, 30)
	aPlayer := audio.NewPlayer(mediaPath)
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	defer func(){
		vDecoder.Stop()
		aPlayer.Stop()
		fmt.Print("\033[?25h") // restores the terminal cursor
	}()

	go func(){
		<- sigChan
		vDecoder.Stop()
		aPlayer.Stop()
		fmt.Print("\033[?25h")
	}()

		if err := vDecoder.Start(); err != nil {
		return fmt.Errorf("video decoder got cooked: %w", err)
	}

	if err := aPlayer.Start(); err != nil {
		return fmt.Errorf("audio player got cooked: %w", err)
	}

	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	for range ticker.C {
		frame, err := vDecoder.NextFrame()
		if err != nil {
			break //video is over
		}

		asciiFrame := render.FrameToString(frame)
		fmt.Print("\033[H]"+ asciiFrame) // moves the cursor to the top-left corner 
	}

	fmt.Println("Playback finished")
	return nil
}