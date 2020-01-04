package main


import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"path"
	"strconv"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	//"image"
	"image/jpeg"
	"log"
	//"time"
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
	if img.Bounds().Dx() > img.Bounds().Dy(){
		h = size
		w = 0
	} else {
		w = size
		h = 0
	}
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio

	resizedImage := resize.Resize(w, h, img, resize.Lanczos3)
	croppedImg, err := cutter.Crop(resizedImage, cutter.Config{
		Width:  int(size),
		Height: int(size),
		//Anchor: image.Point{100, 100},
		Mode:   cutter.Centered, // optional, default value
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

func GenerateForDir(srcFolder, dstFolder string, thumbSize uint){

	fmt.Printf("Enumerating folder %v\n", srcFolder)

	if _, err := os.Stat(dstFolder); os.IsNotExist(err) {
		os.Mkdir(dstFolder, 0666)
	}
	
	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		fmt.Printf("Failed to enumerate folder %v", srcFolder)
		os.Exit(2)
	}

	i := 0
	completion := make(chan bool, 100)
	for _, file := range files {
		if file.IsDir() {
			GenerateForDir(srcFolder + "\\" + file.Name(), dstFolder + "\\" + file.Name(), thumbSize)
		} else {
			go func(f os.FileInfo, index int, done chan<- bool) {

				srcPath := srcFolder + "\\" + f.Name()
				dstPath := dstFolder + "\\" + strconv.Itoa(index) + path.Ext(srcPath)
				thumbPath := dstFolder + "\\" + strconv.Itoa(index) + "_thumb" + path.Ext(srcPath)
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

func main(){
	if len(os.Args) < 4 {
		fmt.Println("Usage: chobi <image folder full path> <out path> <thumb-size>")
		os.Exit(1)
	}

	srcFolder := os.Args[1]
	dstFolder := os.Args[2]
	thumbSize, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Panic("Failed to parse image thumb size")
	}
	
	GenerateForDir(srcFolder, dstFolder, uint(thumbSize))
}
