package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func setAndUsePath(dir string) string {
	exeDir, _ := os.Getwd()
	path := filepath.Join(exeDir, "assets", dir)
	if !CheckPathExists(path) {
		fmt.Printf("Error!! %v does not exist", path)
		os.Exit(3)
	}

	return path
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

	err = SetupAssets("scripts", dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	err = SetupAssets("css", dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	err = SetupAssets("images", dstFolder)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	GenerateImagesIntoDir(name, srcFolder, dstFolder, uint(thumbSize))

}
