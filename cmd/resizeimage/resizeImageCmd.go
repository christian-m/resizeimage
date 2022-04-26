package main

import (
	. "bitbucket.org/christian-m/resizeimage/internal/resize"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	flagInFolder  = flag.String("in", "./", "Input-Folder")
	flagOutFolder = flag.String("out", "", "Output-Folder")
	flagSize      = flag.String("size", "250x250", "max. Size")
	flagBorder    = flag.Int("border", 0, "Border Width")
)

func main() {
	flag.Parse()
	log.Println("Image Resizer started")
	size, err := parseSize(*flagSize)
	if err != nil {
		log.Printf("Can't parse target size: %v\n", err)
	}
	size.AddBorder(*flagBorder)
	outFolder := *flagSize
	if *flagOutFolder != "" {
		outFolder = *flagOutFolder
	}
	errList := resizeFolderImages(*flagInFolder, outFolder, size)
	if errList != nil {
		log.Printf("%v\n", errList.Error())
		os.Exit(1)
	}
	log.Println("Image Resizer finished")
}

func parseSize(s string) (PicSize, error) {
	var ps PicSize
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return ps, fmt.Errorf("%s not in correct format", s)
	}
	var err error
	ps.Width, err = strconv.Atoi(parts[0])
	if err != nil {
		return ps, fmt.Errorf("parseSize: ps.x: %w", err)
	}
	ps.Height, err = strconv.Atoi(parts[1])
	if err != nil {
		return ps, fmt.Errorf("parseSize: ps.y: %w", err)
	}
	return ps, nil
}
