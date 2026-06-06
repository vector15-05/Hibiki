package render

import (
	"image"
	"image/color"
)

const AsciiChars = " .:-=+*#%@"

func FrameToString(frame image.Image) string {
	bounds := frame.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	
	// preallocating a byte slice for speed
	capacity := (width + 1)* height
	img := make([]byte, 0, capacity)

	for y:= 0; y< height; y++ {
		for x := 0; x < width; x++ {
			gray := color.GrayModel.Convert(frame.At(x,y)).(color.Gray)

			index := int(gray.Y) * (len(AsciiChars) - 1) / 255

			img = append(img, AsciiChars[index])
		}
		img = append(img, '\n')
	}

	return string(img)
}

