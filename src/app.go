package main

import (
	"fmt"

	"github.com/Jardielson-s/work-file/src/utils"
)

func main() {
	//MOCK
	mock := utils.File{
		FileName: "test.txt",
		Text:     "Hello!!",
	}

	res, err := mock.CreateFile(mock)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
