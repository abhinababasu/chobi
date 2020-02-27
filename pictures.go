package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/abhinababasu/facethumbnail"
)

func GenerateImagesIntoDir(name, srcFolder, dstFolder string, thumbSize uint, detectFace bool) (int, error) {

	log.Printf("Enumerating folder %v\n", srcFolder)

	if _, err := os.Stat(dstFolder); os.IsNotExist(err) {
		os.Mkdir(dstFolder, 0666)
	}

	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		return 0, fmt.Errorf("Failed to enumerate folder %v", srcFolder)
	}

	dstFolder = filepath.Join(dstFolder, name)
	CreateDirIfNotExist(dstFolder)

	var fd facethumbnail.FaceDetector

	if detectFace {
		pwd, _ := os.Getwd()
		cascadeFile := path.Join(pwd, "facefinder")
		if _, err := os.Stat(cascadeFile); err != nil {
			return 0, fmt.Errorf("Cascade file not found for face detection")
		}

		fd = facethumbnail.GetFaceDetector(cascadeFile)
		fd.Init(-1, -1)
	}

	i := 0
	var wg sync.WaitGroup

	for _, file := range files {
		if file.IsDir() {
			log.Printf("Sub-dir not supported, skipping %v\n", file.Name())
		} else {
			wg.Add(1)
			go func(index int, filename string) {
				defer wg.Done()
				srcPath := filepath.Join(srcFolder, filename)
				dstPath := filepath.Join(dstFolder, strconv.Itoa(index)+path.Ext(srcPath))
				thumbPath := filepath.Join(dstFolder, strconv.Itoa(index)+"_thumb"+path.Ext(srcPath))
				CopyFile(srcPath, dstPath)
				facethumbnail.ResizeImage(fd, srcPath, thumbPath, thumbSize)
				log.Printf(">>>>> Done %v", dstPath)
			}(i, file.Name())
			i++
		}
	}

	wg.Wait()

	return i, nil
}
