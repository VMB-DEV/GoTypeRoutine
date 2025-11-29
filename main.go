package main

import (
	"fmt"
	"os"

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
	//todo : term.GetSize to use all type of term window

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
