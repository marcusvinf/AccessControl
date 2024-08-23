package terminalactions

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/nfnt/resize"
)

const (
	maxWidth  = 600
	maxHeight = 1200
)

func compressBase64Image(base64Image string) (string, error) {
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("error decoding base64 image: %v", err)
	}

	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return "", fmt.Errorf("error decoding image: %v", err)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newWidth := width
	newHeight := height

	if width > maxWidth || height > maxHeight || height > 2*width {
		if width > maxWidth {
			newWidth = maxWidth
			newHeight = int((float64(maxWidth) / float64(width)) * float64(height))
		}

		if newHeight > maxHeight || newHeight > 2*newWidth {
			newHeight = maxHeight
			newWidth = int((float64(maxHeight) / float64(height)) * float64(width))
		}

		img = resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)
	}

	quality := 85
	buffer := new(bytes.Buffer)

	for {
		buffer.Reset()
		err = jpeg.Encode(buffer, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return "", fmt.Errorf("error encoding JPEG: %v", err)
		}

		if buffer.Len() <= 100*1024 || quality <= 5 {
			break
		}

		quality -= 5
	}

	compressedImageBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return compressedImageBase64, nil
}
