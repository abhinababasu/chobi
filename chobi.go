package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: chobi <name> <image folder full path> <out path> <thumb-size>")
		os.Exit(1)
	}

	i := 1
	name := os.Args[i]
	i++
	srcFolder := os.Args[i]
	i++
	dstFolder := os.Args[i]
	i++
	thumbSize, err := strconv.Atoi(os.Args[i])
	i++
	if err != nil {
		log.Panic("Failed to parse image thumb size - %v", err)
	}

	err = SetupAllAssets(dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	imageCount, err := GenerateImagesIntoDir(name, srcFolder, dstFolder, uint(thumbSize))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	log.Printf("Process %v images", imageCount)

	err = GeneratePage(name, imageCount, dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
}
