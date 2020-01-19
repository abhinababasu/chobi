package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	name := flag.String("name", "", "Name of the album")
	srcFolder := flag.String("src", "", "Source directory with images")
	dstFolder := flag.String("dst", "", "Destination directory where the album will generated")
	thumbSize := flag.Uint("size", 150, "Size of thumbnails")
	face := flag.Bool("face", false, "Detect faces in images for generating thumbnails")
	verbose := flag.Bool("v", true, "Verbose logging mode")

	flag.Parse()

	if *name == "" || *srcFolder == "" || *dstFolder == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	if *face {
		log.Print("Will attempt to detect faces")
	}

	err := SetupAllAssets(*dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	imageCount, err := GenerateImagesIntoDir(*name, *srcFolder, *dstFolder, uint(*thumbSize), *face)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	log.Printf("Process %v images", imageCount)

	err = GeneratePage(*name, imageCount, *dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
}
