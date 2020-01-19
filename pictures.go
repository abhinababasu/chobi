package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"

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

	if (detectFace) {
		pwd, _ := os.Getwd()
		cascadeFile := path.Join(pwd, "facefinder")
		if _, err := os.Stat(cascadeFile); err != nil {
			return 0, fmt.Errorf("Cascade file not found for face detection")
		}
		
		fd = facethumbnail.GetFaceDetector(cascadeFile)
		fd.Init(-1, -1)
	}

	i := 0
	for _, file := range files {
		if file.IsDir() {
			log.Printf("Sub-dir not supported, skipping %v\n", file.Name())
		} else {
			srcPath := filepath.Join(srcFolder, file.Name())
			dstPath := filepath.Join(dstFolder, strconv.Itoa(i)+path.Ext(srcPath))
			thumbPath := filepath.Join(dstFolder, strconv.Itoa(i)+"_thumb"+path.Ext(srcPath))
			log.Printf("Copied to %v\n", dstPath)
			CopyFile(srcPath, dstPath)
			i++

			facethumbnail.ResizeImage(fd, srcPath, thumbPath, thumbSize)
		}
	}

	return i, nil
}
