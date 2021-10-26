package main

import (
	"fmt"
	"os"
	"io"
	"time"
	"bufio"
	"strings"
	"math/rand"
)

func main() {
	k := gen(nil)
	r := bufio.NewReader(os.Stdin)

	for {
		chosen := choose(k)
		fmt.Printf("%s:\t", chosen.Symbol)
		
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			ptrerr("reading line", err)
		}

		trimmed := strings.TrimRight(line, "\n")
		if (trimmed == chosen.Romanji) {
			fmt.Println("very good!")
		} else {
			fmt.Println("no, sad!")
		}
	}
}

func choose(k []Romanization) Romanization {
	rand.Seed(time.Now().UnixNano())
	return k[rand.Intn(len(k))]
}

func ptrerr(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err)
	os.Exit(1)
}
