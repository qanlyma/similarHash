package main

import (
	"demo/readwrite"
	"fmt"
)

func Word(s1, s2 string) uint64 { //s1,s2分别为存的文本路径和读取的文本路径
	hashset := readwrite.Readfile(s1)
	testhash := Wordhash(s2)
	hashset = JudgeWord(testhash, hashset, s1)
	fmt.Println(hashset)
	return testhash
}

func Picture(s1, s2 string) uint64 {
	hashset := readwrite.Readfile(s1)
	testhash := PictureToHash(s2)
	hashset = Judgepicture(testhash, hashset, s1)
	fmt.Println(hashset)
	return testhash
}

func main() {
	// 读取图片
	// Picture("phash.txt", "./testPic/2.jpg")
	// 读取文本
	Word("whash.txt", "./dictionary.txt")
}
