package main

import (
	"demo/distance"
	"demo/readwrite"
	"fmt"
)

func Judgepicture(testhash uint64, hashset []uint64, filename string) []uint64 {
	if len(hashset) == 0 {
		readwrite.Fileappend(filename, testhash)
		hashset = append(hashset, testhash)
		fmt.Println("原来为空，插入第一个")
		return hashset
	}
	dis := distance.MinHamming(testhash, hashset)
	fmt.Println("与现存哈希集的汉明距离最小为：", dis)

	if dis > 5 {
		readwrite.Fileappend(filename, testhash)
		hashset = append(hashset, testhash)
		fmt.Println("大于5，插入哈希集")
		fmt.Printf("simhash:0x%x", testhash)
		return hashset
	} else {
		fmt.Println("存在相似元素，不插入哈希集")
		return hashset
	}
}

func JudgeWord(testhash uint64, hashset []uint64, filename string) []uint64 {
	if len(hashset) == 0 {
		readwrite.Fileappend(filename, testhash)
		hashset = append(hashset, testhash)
		fmt.Println("原来为空，插入第一个")
		return hashset
	}
	dis := distance.MinHamming(testhash, hashset)
	fmt.Println("与现存哈希集的汉明距离最小为：", dis)

	if dis > 10 {
		readwrite.Fileappend(filename, testhash)
		hashset = append(hashset, testhash)
		fmt.Println("大于10，插入哈希集")
		fmt.Printf("simhash:0x%x", testhash)
		return hashset
	} else {
		fmt.Println("存在相似元素，不插入哈希集")
		return hashset
	}
}
