package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"liuma/utils"
	"log"
	"os"
	"testing"
)

func TestFileHash(t *testing.T) {



	filebyte, err := ioutil.ReadFile("/Users/lzl/Documents/头像.jpeg"); if err != nil {
		log.Println(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
	log.Println(len(filebyte))
	log.Println(utils.CalculateHash(bytes.NewReader(filebyte)))






}
