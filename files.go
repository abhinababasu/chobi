package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	assetFolder      = "assets"
	galleryNameToken = "GALLERY_NAME"
	imageCountToken  = "IMG_COUNT"
)

// subfolders in assets that needs to be copied to target
var assets = []string{"scripts", "css", "images"}

// CreateDirIfNotExist creates all dirs in the path if does not exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// CheckPathExists checks if a file path exists on the filesystem
func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// CopyDir copies a directory from source to destination recursively
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// CopyFile copies a single file from src to dst
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

// SetupAssets copies a named asset to the destination
func SetupAssets(name, destination string) error {
	log.Printf("Setting up asset '%v' in %v", name, destination)
	exeDir, _ := os.Getwd()
	srcPath := filepath.Join(exeDir, assetFolder, name)
	if !CheckPathExists(srcPath) {
		return fmt.Errorf("Asset %v does not exist", srcPath)
	}

	dstPath := path.Join(destination, name)
	err := CopyDir(srcPath, dstPath)
	return err

}

// SetupAllAssets copies all assets to the destination
func SetupAllAssets(destination string) error {
	for _, asset := range assets {
		err := SetupAssets(asset, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

// GeneratePage generates the main html page into the destination
func GeneratePage(name string, imageCount int, destination string) error {
	exeDir, _ := os.Getwd()
	templateFilePath := filepath.Join(exeDir, assetFolder, "GALLERY_NAME.htm")
	if !CheckPathExists(templateFilePath) {
		return fmt.Errorf("Template %v does not exist", templateFilePath)
	}

	buffer, err := ioutil.ReadFile(templateFilePath)
	srcStr := string(buffer)
	if err != nil {
		return err
	}

	finalStr := strings.ReplaceAll(srcStr, galleryNameToken, name)
	finalStr = strings.ReplaceAll(finalStr, imageCountToken, strconv.Itoa(imageCount))
	dstFile := filepath.Join(destination, name+".html")

	ioutil.WriteFile(dstFile, []byte(finalStr), 0644)

	return nil
}
