package main

import (
	"fmt"
	"os"

	"github.com/vector15-05/Hibiki/engine"
)

func main(){
	//videoFile := "cmd/example/bad_apple.mp4"
	videoFile := "Enter your video file name here"

	if _,err := os.Stat(videoFile); os.IsNotExist(err) {
		fmt.Printf("Error: Video file %s not found\n", videoFile)
		fmt.Println("Please make sure the video file is in the same directory as the executable.")
		os.Exit(1)
	}

	fmt.Println("Starting playback... Press Ctrl+C to stop.")

	if err := engine.Run(videoFile); err != nil {
		fmt.Printf("Error during playback: %v\n", err)
		os.Exit(1)
	}
}