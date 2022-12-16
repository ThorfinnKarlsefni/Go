package chapter12

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FileCopy(){
	source,_ := filepath.Abs("../source.txt")
	fmt.Println(source)
	CopyFile("./target.txt","./source.txt")
	fmt.Println("Copy done!")
}

func CopyFile (dstName,srcName string) (written int64,err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(err,1)
		return
	}

	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err,2)
		return
	}

	defer dst.Close()
	return io.Copy(dst,src)
}