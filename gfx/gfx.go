package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("title.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	m, _ := img.(*image.Gray)
	var sprites []byte
	for sy := 0; sy < m.Bounds().Dy(); sy += 8 {
		for sx := 0; sx < m.Bounds().Dx(); sx += 8 {
			var sprite [16]byte
			for dy := 0; dy < 8; dy++ {
				var a, b byte
				for dx := 0; dx < 8; dx++ {
					a <<= 1
					b <<= 1
					switch c := m.GrayAt(sx+dx, sy+dy).Y; c {
					case 0:
						a |= 0x1
					case 25:
						a |= 0x1
						b |= 0x1
					case 123, 165:
						b |= 0x1
					case 255:
					default:
						log.Fatalf("unknown gray %d", c)
					}
				}
				sprite[dy] = a
				sprite[dy+8] = b
			}
			sprites = append(sprites, sprite[:]...)
		}
	}
	fmt.Printf("@title-image %.2x %.2x ( w, h )",
		m.Bounds().Dx()/8, m.Bounds().Dy()/8)
	for i, b := range sprites {
		if i%0x10 == 0 {
			fmt.Printf("\n\t")
		} else {
			fmt.Printf(" ")
		}
		fmt.Printf("%.2x", b)
	}
	fmt.Println()
	log.Printf("%d sprite bytes", len(sprites))
}
