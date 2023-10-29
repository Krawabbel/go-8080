package main

import (
	"os"

	"github.com/Krawabbel/go-8080/cpm"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		if err := cpm.Run(os.Args[i]); err != nil {
			panic(err)
		}
	}
}
