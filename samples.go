package main

// sox sfx.wav -t raw -c 1 -r 44100 -b 8 -e unsigned-integer sfx.raw

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("@sfx")
	for _, fn := range os.Args[1:] {
		b, err := os.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\t&%s %.4x", strings.TrimSuffix(fn, ".raw"), uint16(len(b)))
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
