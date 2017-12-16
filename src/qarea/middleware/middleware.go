package middleware

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"path/filepath"
	"regexp"
	"strings"
)

var Shared string

const directory string = "./"

func FindWord(keyword string) {

	dir, err := os.Open(directory)
	if err != nil {
		fmt.Println(err)
	}
	defer dir.Close()

	filez, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	defer dir.Close()

	var reg = regexp.MustCompile(`(?sm)\<(\/?)script.*?\/script>|(?sm)\<(\/?)a.*?\/a>|(?sm)\<(\/?)svg.*?\/svg>|\<(\/?)link.*?\/>|(?sm)\<(\/?)ul.*?\/ul>|(?sm)\<(\/?)li.*?\/li>|\<(\/?)meta.*?\/>|(?sm)\<(\/?)style.*?\/style>|(?sm)\<(\/?)form.*?\/form>`)

	match := make(map[string]int)

	for _, file := range filez {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".txt" {
				fileName := file.Name()

				fl, err := os.Open(fileName)
				if err != nil {
					log.Fatal(err)
				}
				defer fl.Close()

				scanner := bufio.NewScanner(fl)

				outArr := make([]string, 0)
				countOfWords := 0

				for scanner.Scan() {
					outHTML := scanner.Text()
					outArr = append(outArr, outHTML)
				}
				str := strings.Join(outArr[:], "")
				outReg := reg.ReplaceAllString(str, " ")
				countOfWords += strings.Count(outReg, keyword)
				match[fileName] = countOfWords

			}
		}
	}
	fmt.Println("\n",match,"\n")
	maxOfMap(match)
}

func maxOfMap(finalMatch map[string]int) {
	max := 0
	fileKey := ""
	outValue := 0
	for key, value := range finalMatch {
		if value > max {
			max = value
			if value == max {
				fileKey = key
				outValue = value
			}
		}
	}
	fmt.Println("Priority article is", fileKey, "which have", outValue, "matches.")
	getText(fileKey)
}

func getText(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var regOut = regexp.MustCompile(`(?sm)\<(\/?)script.*?\/script>|(?sm)\<(\/?)a.*?\/a>|(?sm)\<(\/?)svg.*?\/svg>|\<(\/?)link.*?\/>|(?sm)\<(\/?)ul.*?\/ul>|(?sm)\<(\/?)li.*?\/li>|\<(\/?)meta.*?\/>|(?sm)\<(\/?)style.*?\/style>|(?sm)\<(\/?)form.*?\/form>|\<img.*?\/>|\<(\/?)noscript.*?\/>.*\<\/noscript\>`)

	scanner := bufio.NewScanner(file)

	outArr := make([]string, 0)

	for scanner.Scan() {
		outHTML := scanner.Text()
		outArr = append(outArr, outHTML)
	}
	str := strings.Join(outArr[:], "")
	outReg := regOut.ReplaceAllString(str, " ")

	Shared = outReg
}