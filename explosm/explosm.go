package explosm

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/mmcdole/gofeed"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/log"
)

const (
	explosmFeedURL = "http://feeds.feedburner.com/Explosm"
)

var (
	imgRegexp = regexp.MustCompile(`(?s)<div id="comic-container">\s*?<div class="row">\s*?<div class="small-12 medium-12 large-12 columns">(.*?)<\/div>`)
)

type Explosm struct {
	RefreshInterval time.Duration

	data    []byte
	rssData channel
}

func (e *Explosm) Run(abort chan struct{}) {
	// Execute it the first time
	e.Do(nil)
	for {
		select {
		case <-abort:
			return
		case <-time.After(e.RefreshInterval):
			e.Do(nil)
		}
	}
}

func (e *Explosm) Do(r *http.Request) error {
	fp := gofeed.NewParser()
	if r != nil {
		fp.Client = urlfetch.Client(appengine.NewContext(r))
	}
	feed, err := fp.ParseURL(explosmFeedURL)
	if err != nil {
		return err
	}
	var is []Item
	for _, i := range feed.Items {
		e := Explosm{}
		if err := e.GetData(i.Link, r); err != nil {
			safeLog(r, "%s", err)
		}
		imgEle := FindComicURL(e.data)
		if imgEle == "" {
			safeLog(r, "cannot find image. Data is: %s", e.data)
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

func (e *Explosm) GetData(url string, r *http.Request) error {
	d, err := e.getDataFromNet(url, r)
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

func (e *Explosm) getDataFromNet(url string, r *http.Request) ([]byte, error) {
	httpClient := http.DefaultClient
	if r != nil {
		httpClient = urlfetch.Client(appengine.NewContext(r))
	}
	res, err := httpClient.Get(url)
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

// safeLog only logs if request is not nil
func safeLog(r *http.Request, format string, args ...interface{}) {
	if r == nil {
		return
	}
	log.Errorf(appengine.NewContext(r), format, args)
}
