package main

// sox sfx.wav -t raw -c 1 -r 44100 -b 8 -e unsigned-integer sfx.raw

import (
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("sfx.raw")
	if err != nil {
		log.Fatal(err)
	}
	var samples [][]byte
	for len(b) > 0 {
		quiet := 0
		for i := range b {
			//log.Printf("%.2x %d", b[i], quiet)
			switch b[i] {
			case 0x7f, 0x80, 0x81:
				quiet++
			default:
				quiet = 0
			}
			if quiet > 20 {
				if i < 256 {
					log.Printf("found a sample %d bytes long (discarding)", i)
				} else {
					log.Printf("found a sample %d bytes long", i)
					samples = append(samples, b[:i])
				}
				b = b[i:]
				break
			}
		}
		found := false
		for i := range b {
			switch b[i] {
			case 0x7e, 0x7f, 0x80, 0x81, 0x82:
			default:
				log.Printf("found %.2x", b[i])
				found = true
			}
			if found {
				b = b[i:]
				break
			}
		}
		if !found {
			log.Printf("found no more samples")
			break
		}
	}
	fmt.Println("@samples")
	for i, b := range samples {
		fmt.Printf("\t&s%d %.4x", i, uint16(len(b)))
		for i, b := range b {
			if i%0xf == 0 {
				fmt.Printf("\n\t\t")
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf("%.2x", b)
		}
		fmt.Println()
	}
}
