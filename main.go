package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/lunny/html2md"
)

var (
	baseURL      = "https://blog.scottlogic.com"
	blogLinks    []string
	blogImgLinks []string
)

func main() {
	getLinks()

	fmt.Println("Got Blog Links")
	for _, link := range blogLinks {
		fmt.Println("Getting content for blog post: ", link)

		postMrkDwn := getPostContent(link)
		blogFilePath := fmt.Sprintf("blogs/%s", fileNameFromLink(link))

		err := ioutil.WriteFile(blogFilePath, postMrkDwn, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Got all blogs content, Getting Images...")

	for _, img := range blogImgLinks {
		downloadImg(img)
	}

	fmt.Println("Done")
}

func getLinks() {
	doc, _ := htmlquery.LoadURL(fmt.Sprintf("%s/rhamilton/", baseURL))

	blogLinkNodes := htmlquery.Find(doc, "//div[@class='title']/a")

	for _, link := range blogLinkNodes {
		blogLinks = append(blogLinks, htmlquery.SelectAttr(link, "href"))
	}
}

func getPostContent(postLink string) []byte {
	doc, _ := htmlquery.LoadURL(fmt.Sprintf("%s%s", baseURL, postLink))

	contentNode := htmlquery.Find(doc, "//div[@class='post-content cell large-6']/div[@class='cell']")
	postImgs := htmlquery.Find(contentNode[0], "//img")

	for _, link := range postImgs {
		blogImgLinks = append(blogImgLinks, htmlquery.SelectAttr(link, "src"))
	}

	postHTML := strings.Replace(htmlquery.OutputHTML(contentNode[0], false), "/rhamilton/assets/", "assets/", -1)
	md := html2md.Convert(postHTML)

	return []byte(md)
}

func downloadImg(link string) {
	url := fmt.Sprintf("%s%s", baseURL, link)
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		fmt.Println(e)
	}
	defer response.Body.Close()

	//open a file for writing
	imgFilePath := fmt.Sprintf("assets/%s", fileNameFromLink(link))
	file, err := os.Create(imgFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
	}
}

func fileNameFromLink(link string) string {
	ss := strings.Split(link, "/")
	return strings.Replace(ss[len(ss)-1], ".html", ".md", 1)
}
