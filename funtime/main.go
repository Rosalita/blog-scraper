package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"

	"github.com/lunny/html2md"
)

var (
	baseURL      = "https://blog.scottlogic.com"
	blogLinks    []string
	blogImgLinks []string
)

func main() {

	files, err := ioutil.ReadDir("html")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		f, err := os.Open("html/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()

		html, err := ioutil.ReadAll(f)
		markdown := html2md.Convert(string(html))

		fmt.Println("****")
		fmt.Println(markdown)

		fullname := strings.Split(file.Name(), ".")
		baseName := fullname[0]
		markdownName := baseName + ".md"

		outputPath := fmt.Sprintf("md/%s", markdownName)

		err = ioutil.WriteFile(outputPath, []byte(markdown), 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
