package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

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

	GenerateImagesIntoDir(name, srcFolder, dstFolder, uint(thumbSize))
}
