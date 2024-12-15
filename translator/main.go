package main

import (
	"fmt"
	"flag"
	// "io/ioutil"
	"strings"
	"os"
	"sync"

	"github.com/puttarajkoliwad/go_projects/translator/cli"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

func main() {
	flag.Parse()
	
	strChan := make(chan string)
	
	wg.Add(1)

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	reqBody := &cli.RequestBody{
		sourceLang,
		targetLang,
		sourceText,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg)

	processedStr := strings.ReplaceAll(<-strChan, "+", " ")
	// fmt.Println(processedStr[0])
	fmt.Println("\n", len([]rune(processedStr)))

	close(strChan)
	wg.Wait()
	os.Exit(0)
}