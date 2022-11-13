package readwrite

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件
func Readfile(filename string) []uint64 {
	file, err := os.Open(filename)
	var hashset []uint64
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			if Str2Uint64(line) != 0 {
				hashset = append(hashset, Str2Uint64(line))
			}
			break
		} else {
			hashset = append(hashset, Str2Uint64(line))
		}
	}
	return hashset
}

func Str2Uint64(str string) uint64 {
	var res uint64
	fmt.Sscanf(str, "0x%x", &res)
	return res
}

func Fileappend(filename string, hash uint64) {

	Writefile(filename, fmt.Sprintf("0x%016x", hash))

}

// 写文件（追加）
func Writefile(filename string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content+"\n"), n)
		if err != nil {
			fmt.Println("cacheFileList.yml file write failed. err: " + err.Error())
		}
	}
	defer f.Close()
	return err
}
