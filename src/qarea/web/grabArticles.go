package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"io"

	"github.com/PuerkitoBio/goquery"
)

const urlPage = "https://habrahabr.ru/all/page1/"

/*const directory string = ".." + string(filepath.Separator) + "articles" + string(filepath.Separator)*/

const directory string = "./"

func GrabArticles() {

	doc, err := goquery.NewDocument(urlPage)
	if err != nil {
		log.Fatal(err)
	}

	arr := make([]string, 0, 10)

	doc.Find(".post__title").Each(func(index int, item *goquery.Selection) {
		linkTag := item.Find(".post__title_link")
		link, _ := linkTag.Attr("href")

		fmt.Println(link)

		arr = append(arr, link)
	})

	for index := 0; index < len(arr); index++ {
		str := arr[index]
		resp, err := http.Get(str)
		if err != nil {
			fmt.Println("error")
		}
		defer resp.Body.Close()

		numOfDoc := strconv.Itoa(index)

		file, err := os.Create(directory + "Article" + numOfDoc + ".txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		io.Copy(file, resp.Body)
	}
}