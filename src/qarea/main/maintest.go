package main

import (
	"os"
	"fmt"
	"log"
/*	"io/ioutil"*/
	/*"strings"*/

/*	"github.com/kennygrant/sanitize"*/
	"github.com/PuerkitoBio/goquery"
/*	"qarea/web"
	"qarea/middleWareTest"*/
)

func main() {
	Grab()
}

func Grab() {
	f, e := os.Open("Article0.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f); 
	if err != nil { 
		log.Fatal(err)
	}



	doc.Find(".post_full").Each(func(index int, item *goquery.Selection) {
		i := item.Text()
		fmt.Println(i)
	})
}