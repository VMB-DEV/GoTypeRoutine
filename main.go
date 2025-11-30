package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Failed to restore terminal:", err)
		}
	}()

	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		width, height = 0, 0
	}

	fmt.Printf("term: %dw%dh\r\n", width, height)
	fmt.Printf("%s\r\n\n", strings.Repeat("_", width-1))
	fmt.Print("Just Type (ESC to quit):\r")
	//todo: care about arrow keys later since it take 3 byte
	buf := make([]byte, 1)

	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			break
		}

		key := buf[0]

		// ESC key is ASCII 27
		if key == 27 {
			fmt.Print("\n\ryolo\n\r")
			break
		}

		fmt.Printf("%c", key)
	}
}

//todo: make a struc for Ansii handling
//todo: make a struc for Ascii handling
