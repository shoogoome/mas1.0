package test

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"liuma/utils"
	"os"
	"testing"
)

func TestGzip(t *testing.T) {
	file, _ := os.Open("/Users/lzl/Documents/project/liuma/liuma/tests/123.txt")

	//info, _ := file.Stat()
	fileByte, _ := ioutil.ReadAll(file)
	//filegzip, _ := utils.GzipFile(fileByte, "123.txt", info.Size(), "thisishash")
	filegzip, _ := utils.GzipFile(fileByte, "123.txt")
	filegnzip, _ := utils.GunzipFile(filegzip)
	_ = ioutil.WriteFile("/Users/lzl/Documents/project/liuma/liuma/123.txt", filegnzip, 0666)

}

func TestUnGzip(t *testing.T) {
	srcFile, err := os.Open("/Users/lzl/Documents/project/liuma/liuma/tests/qwe.tar.gz")
	if err != nil {
		fmt.Sprintf("%v", err)
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		fmt.Sprintf("%v", err)
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	hdr, err := tr.Next()
	fmt.Println(hdr.Name)
	if err != nil {
		fmt.Sprintf("%v", err)
	}
	filename := "/Users/lzl/Documents/project/liuma/liuma/tests/" + hdr.Name
	file, err := os.Create(filename)
	if err != nil {
		fmt.Sprintf("%v", err)
	}
	io.Copy(file, tr)
}