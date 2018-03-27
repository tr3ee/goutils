package AsciiArt

import (
	"bytes"
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

// Image2AsciiArt convert an image into a specified size character picture using chars as a dictionary
func Image2AsciiArt(img image.Image, chars []byte, width, height int) []byte {
	img = imaging.Resize(img, width, height, imaging.Lanczos)
	rect := img.Bounds()
	buf := new(bytes.Buffer)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			var c uint8
			if len(chars) <= 256 {
				g := ColorToGray(img.At(x, y))
				c = GrayToChar(g, chars)
			} else {
				g := ColorToGray16(img.At(x, y))
				c = Gray16ToChar(g, chars)
			}
			buf.WriteByte(c)
		}
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func ColorToGray(c color.Color) uint8 {
	return color.GrayModel.Convert(c).(color.Gray).Y
}

func GrayToChar(gray uint8, chars []byte) byte {
	pos := int(gray) * len(chars) >> 8
	return chars[pos]
}

func ColorToGray16(c color.Color) uint16 {
	return color.Gray16Model.Convert(c).(color.Gray16).Y
}

func Gray16ToChar(gray uint16, chars []byte) byte {
	pos := int(gray) * len(chars) >> 16
	return chars[pos]
}
