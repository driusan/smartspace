package main

import "os"
import "io/ioutil"
import "fmt"
import "unicode/utf8"
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: smartspace filename\n");
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s for reading\n", os.Args[1])
		return 
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var lastRune rune
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r == ' ' {
			switch lastRune {
				case '.':
					// put an em-space at the start of a sentence
					fmt.Printf("\u2001")
				case ',', ';', ':', '!', '?':
					// put an en-space after any non-sentence punctuation
					fmt.Printf("\u2002")
				default: 
					// use 1/3 em.
					fmt.Printf("\u2004")
			}
		} else {
			fmt.Printf("%c", r)
		}
		lastRune = r
		b = b[size:]
	}
}