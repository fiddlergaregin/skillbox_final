package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var path_to_external_data string = "../garegin_go_addon/"

func main() {
	f, err := os.OpenFile(path_to_external_data+"sms.data", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file_content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file_content))
}
