package main

import (
	. "bitbucket.org/christian-m/resizeimage/internal/resize"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type resizeArgs struct {
	inPath    string
	outPath   string
	size      PicSize
	addBorder bool
}

func resizeFolderImages(inFolder, outFolder string, size PicSize) *errorList {
	errList := &errorList{}
	err := os.MkdirAll(outFolder, 0755)
	if err != nil {
		errList.add(fmt.Errorf("can't create target path: %v", err))
		return errList
	}
	dir, err := ioutil.ReadDir(inFolder)
	if err != nil {
		errList.add(fmt.Errorf("can't read from folder %s: %v", inFolder, err))
		return errList
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error)
	resizeChan := make(chan resizeArgs)
	wg.Add(5)
	go worker(wg, resizeChan, errChan)
	go worker(wg, resizeChan, errChan)
	go worker(wg, resizeChan, errChan)
	go worker(wg, resizeChan, errChan)
	go worker(wg, resizeChan, errChan)
	go func(errList *errorList, errChan chan error) {
		for err := range errChan {
			errList.add(err)
		}
	}(errList, errChan)

	for _, fi := range dir {
		if fi.IsDir() || !useFile(fi.Name()) {
			log.Printf("Omit file %q, directory or wrong file format", fi.Name())
			continue
		}
		inPath := filepath.Join(inFolder, fi.Name())
		outPath := filepath.Join(outFolder, fi.Name())
		resizeChan <- resizeArgs{inPath, outPath, size, false}
	}
	close(resizeChan)
	close(errChan)
	wg.Wait()
	if errList.hasError() {
		return errList
	}
	return nil
}

func useFile(s string) bool {
	allowed := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(s)
	for _, e := range allowed {
		if strings.EqualFold(ext, e) {
			return true
		}
	}
	return false
}

func worker(wg *sync.WaitGroup, c chan resizeArgs, errChan chan error) {
	for a := range c {
		if a.addBorder {
			log.Println("Resize file and add border", a.inPath)
		} else {
			log.Println("Resize file", a.inPath)
		}
		inFile, err := os.Open(a.inPath)
		if err != nil {
			errChan <- fmt.Errorf("failed open file %s: %v", a.inPath, err)
			continue
		}
		outFile, err := os.OpenFile(a.outPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			errChan <- fmt.Errorf("failed creating file %s: %v", a.outPath, err)
			inFile.Close()
			continue
		}
		img, f, err := image.Decode(inFile)
		if err != nil {
			errChan <- fmt.Errorf("error decoding image: %w", err)
		}
		err = Resize(a.size, f, img, outFile)
		if err != nil {
			errChan <- fmt.Errorf("failed resizing image %s: %v", a.inPath, err)
		}
	}
	wg.Done()
}
