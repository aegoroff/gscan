package main

import (
	"fmt"
	"github.com/aegoroff/dirstat/scan"
	"os"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	h := &fh{}
	start := time.Now()
	scan.Scan(argsWithoutProg[0], scan.NewOsFs(), h)
	elapsed := time.Since(start)
	fmt.Printf("files: %d folders: %d elapsed: %v\n", h.files, h.folders, elapsed)
}

type fh struct {
	files   int64
	folders int64
}

func (h *fh) Handle(evt *scan.Event) {
	if evt.File != nil {
		h.files += 1
	}
	if evt.Folder != nil {
		h.folders += 1
	}
}
