package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/corona10/goimagehash"
	"github.com/yanyiwu/gosimhash"
)

func PictureToHash(fPATH string) uint64 {
	file, _ := os.Open(fPATH)
	defer file.Close()

	img, _, _ := image.Decode(file)
	hash, _ := goimagehash.PerceptionHash(img)

	return hash.GetHash()
}

func Wordhash(fPATH string) uint64 {
	hasher := gosimhash.New("./dict/jieba.dict.utf8", "./dict/hmm_model.utf8", "./dict/idf.utf8", "./dict/stop_words.utf8")
	defer hasher.Free()
	file, err := os.Open(fPATH)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	file.Read(buffer)
	if err != nil {
		fmt.Println(err)

	}

	return hasher.MakeSimhash(string(buffer), 5)
}
