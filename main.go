package Google

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gocolly/colly"
)

type ImageStruct struct {
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Source string `json:"source"`
	Image  string `json:"image"`
}

func GetImages(keyword string) ([]ImageStruct, error) {
	// url := fmt.Sprintf("https://google.com/search?q=%s&tbm=isch", url.QueryEscape(keyword))
	url := fmt.Sprintf("https://google-scrapper.orewa.workers.dev/?img=%v", url.QueryEscape(keyword))

	c := colly.NewCollector(
	// colly.UserAgent("User-Agent: Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36"),
	)

	// set variable
	var Images []ImageStruct

	// grab image
	c.OnHTML(".islrtb.isv-r", func(e *colly.HTMLElement) {
		h, _ := strconv.Atoi(e.Attr("data-oh"))
		w, _ := strconv.Atoi(e.Attr("data-ow"))
		result := ImageStruct{
			Image:  e.Attr("data-ou"),
			Height: h,
			Weight: w,
			Source: e.Attr("data-st"),
		}
		Images = append(Images, result)
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url) // start

	if len(Images) == 0 {
		return nil, fmt.Errorf("we cant find images")
	}

	return Images, nil
}
