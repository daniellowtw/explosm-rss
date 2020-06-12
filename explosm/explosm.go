package explosm

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
	"log"

	"github.com/mmcdole/gofeed"
)

const (
	explosmFeedURL = "http://feeds.feedburner.com/Explosm"
)

var (
	// Compare the test and source on the website to see if this regex is still valid
	imgRegexp = regexp.MustCompile(`(?s)<div id="comic-wrap">(.*?)</div>`)
)

type Explosm struct {
	RefreshInterval time.Duration

	data    []byte
	rssData channel
}

func (e *Explosm) Run(abort chan struct{}) {
	// Execute it the first time
	if err := e.Do(); err != nil {
		log.Printf("could not run: %s", err)
	}
	for {
		select {
		case <-abort:
			return
		case <-time.After(e.RefreshInterval):
			if err := e.Do(); err != nil {
				log.Printf("could not run: %s", err)
			}
		}
	}
}

func (e *Explosm) Do() error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(explosmFeedURL)
	if err != nil {
		return err
	}
	var is []Item
	for _, i := range feed.Items {
		e := Explosm{}
		if err := e.GetData(i.Link); err != nil {
			// Tolerate bad entries and just log them
			log.Printf("could not get data from link: %s\n", err)
			continue
		}
		imgEle := FindComicURL(e.data)
		if imgEle == "" {
			// Tolerate bad entries and just log them
			log.Printf("could not find image link from feed item\n")
		}
		is = append(is, Item{
			Title:       i.Title,
			Link:        i.Link,
			Description: CDataTest{imgEle},
			Category:    i.Categories,
			Guid:        i.GUID,
			PubDate:     i.Published,
		})
	}
	e.rssData = channel{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		Item:        is,
		Image: Image{
			URL:   "//files.explosm.net/img/favicons/site/favicon-96x96.png",
			Link:  feed.Link,
			Title: feed.Title,
		},
	}
	return nil
}

func (e *Explosm) GetData(url string) error {
	d, err := e.getDataFromNet(url)
	if err != nil {
		return err
	}
	e.data = d
	return nil
}

func FindComicURL(data []byte) string {
	matches := imgRegexp.FindSubmatch(data)
	if len(matches) > 1 {
		return string(matches[1])
	}
	return ""
}

func (e *Explosm) getDataFromNet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *Explosm) Generate() string {
	if len(e.rssData.Item) == 0 {
		return "Please try again. There was an error retrieving the feeds: no feeds."
	}
	return generate(e.rssData)
}
