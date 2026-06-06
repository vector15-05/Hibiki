package render

const AsciiPalette = " .:-=+*#%@"

func FrameToString(frameData []byte, width, height int) string {
	// preallocating a byte slice for speed
	capacity := (width + 1)* height
	img := make([]byte, 0, capacity)

	i := 0
	for y:= 0; y< height; y++ {
		for x := 0; x < width; x++ {
			pixelLuminance := int(frameData[i])
			index := pixelLuminance * (len(AsciiPalette) - 1) / 255
			img = append(img, AsciiPalette[index])
			i++
		}
		img = append(img, '\n')
	}

	return string(img)
}

