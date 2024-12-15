package cli

import(
	"fmt"
	"net/http"
	"log"
	"sync"
	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translationURL = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, strChan chan string, wg *sync.WaitGroup) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", translationURL, nil)

	if err != nil {
		log.Fatal("Error creating request %s", err)
		return
	}
	
	query := req.URL.Query()

	query.Add("client", "gtx")

	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("Error contacting gxt: %s", err)
	}

	// defer: close the body once this entire function is run (like a callback)
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		strChan <- "You have been rate limited. Please try again!"
		wg.Done()
		return
	}

	parsedJSON, err := gabs.ParseJSONBuffer(res.Body)

	if err != nil {
		log.Fatal("Error parsing JSON", err)
	}

	nestOne, err := parsedJSON.ArrayElement(0)

	if err != nil {
		log.Fatal("Error fetching 0th element", err)
	}

	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatal("Error fetching 0th nested element", err)	
	}

	translatedStr, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatal("error fetching translated string", err)
	}
	fmt.Println("cli", translatedStr)
	strChan <- translatedStr.Data().(string)
	wg.Done()
}