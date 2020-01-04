package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func ResizeImage(srcPath, dstPath string, size uint) {
	file, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	w := uint(0)
	h := uint(0)
	if img.Bounds().Dx() > img.Bounds().Dy() {
		h = size
		w = 0
	} else {
		w = size
		h = 0
	}

	// TODO: Support portrait thumbnail that uses top part of photo
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio

	resizedImage := resize.Resize(w, h, img, resize.Lanczos3)
	croppedImg, err := cutter.Crop(resizedImage, cutter.Config{
		Width:  int(size),
		Height: int(size),
		//Anchor: image.Point{100, 100},
		Mode: cutter.Centered, // optional, default value
	})

	out, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, croppedImg, nil)
	log.Printf("Generated %v", dstPath)
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func GenerateImagesIntoDir(name, srcFolder, dstFolder string, thumbSize uint) {

	log.Printf("Enumerating folder %v\n", srcFolder)

	if _, err := os.Stat(dstFolder); os.IsNotExist(err) {
		os.Mkdir(dstFolder, 0666)
	}

	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		fmt.Printf("Failed to enumerate folder %v", srcFolder)
		os.Exit(2)
	}

	dstFolder = filepath.Join(dstFolder, name)
	CreateDirIfNotExist(dstFolder)

	i := 0
	completion := make(chan bool, 100)
	for _, file := range files {
		if file.IsDir() {
			log.Printf("Sub-dir not supported, skipping %v\n", file.Name())
		} else {
			go func(f os.FileInfo, index int, done chan<- bool) {

				srcPath := filepath.Join(srcFolder, f.Name())
				dstPath := filepath.Join(dstFolder, strconv.Itoa(index)+path.Ext(srcPath))
				thumbPath := filepath.Join(dstFolder, strconv.Itoa(index)+"_thumb"+path.Ext(srcPath))
				log.Printf("Copied to %v\n", dstPath)
				CopyFile(srcPath, dstPath)
				ResizeImage(srcPath, thumbPath, thumbSize)
				done <- true
			}(file, i, completion)
			i++
		}
	}

	for j := 0; j < i; j++ {
		select {
		case <-completion:
		}
	}
}
