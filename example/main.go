package main

import (
		"os"
		"log"
		"io/ioutil"
		"github.com/wangtiga/exterror"
)

func main() {
		err := SaveFile("tmp.txt", []byte("hello exterror"))
		if nil != err  {
				log.Printf("SaveFile Fail! Cause: %s \n", err.Error())
		}
}

func SaveFile(name string, data []byte) (error) {
	tmpDir := os.TempDir()
	tmpfile, err := ioutil.TempFile(tmpDir, name)
	if nil!= err {
		return exterror.New(err)
	}

	if _, err := tmpfile.Write(data); err != nil {
		return exterror.New(err)
	}

	if err := tmpfile.Close(); err != nil {
		return exterror.New(err)
	}

	log.Printf("SaveFile Succ! Paht: %s \n", tmpfile.Name())
	return nil
}


