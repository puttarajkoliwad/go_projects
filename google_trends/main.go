// Why google-trends is a great project? Ans: https://www.oberlo.com/blog/google-trends

package main

import (
	"fmt"
	"net/http"
	"encoding/xml"
	"os"
	// "time"
	"io/ioutil"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel	`xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title string	`xml:"title"`
	Traffic string	`xml:"approx_traffic"`
	Description string	`xml:"description"`
	Link string	`xml:"link"`
	// PublishDate time.Time	`xml:"pubDate"`
	PictureLink string	`xml:"picture"`
	PictureSource string	`xml:"picture_source"`
	NewsItems []News 		`xml:"news_item"`
}

type News struct {
	HeadLine string `xml:"news_item_title"`
	Link string	`xml:"news_item_url"`
	Picture string `xml:"news_item_picture"`
	Source string	`xml:"news_item_source"`
}

func main() {
	var r RSS

	data := readGoogleTrends()

	if err := xml.Unmarshal(data, &r); err != nil {
		fmt.Println("Error reading xml data:", err)
		os.Exit(1)
	}

	fmt.Println(r.Channel.Title)
	for i, item := range r.Channel.Items {
		if i == 3 {
			break
		}
		fmt.Println(i+1, item.Title)
		fmt.Println("  Link:", item.Link)
		fmt.Println("  Picture link:", item.PictureLink)
		fmt.Println("  Picture source:", item.PictureSource)
		fmt.Println("  News items:\n")
		
		for j, news := range item.NewsItems {
			if j == 3 {
				break
			}
			fmt.Printf("  %d.%d) %s\n", i+1, j+1, news.HeadLine)
			fmt.Println("       Link:", news.Link)
			fmt.Println("       Picture:", news.Picture)
			fmt.Println("       Source:", news.Source, "\n")
		}

		fmt.Println("\n\n");
	}
}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trending/rss?geo=IN")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(resp.Status)
	return resp
}