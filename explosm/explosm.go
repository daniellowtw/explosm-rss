package explosm

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/mmcdole/gofeed"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

const (
	explosmFeedURL = "http://feeds.feedburner.com/Explosm"
)

var (
	r = regexp.MustCompile(`(?s)<div id="comic-container">
<div class="row">
<div class="small-12 medium-12 large-12 columns">
(.*?)
</div>
</div>
</div>`)
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
		client := urlfetch.Client(appengine.NewContext(r))
		fp.Client = client
	}
	feed, err := fp.ParseURL(explosmFeedURL)
	if err != nil {
		println(err.Error())
		return err
	}
	var is []Item
	for _, i := range feed.Items {
		e := Explosm{}
		if err := e.GetData(i.Link, r); err != nil {
			log.Fatal(err)
		}
		is = append(is, Item{
			Title:       i.Title,
			Link:        i.Link,
			Description: CDataTest{e.FindComicURL()},
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

func (e *Explosm) FindComicURL() string {
	matches := r.FindSubmatch(e.data)
	if len(matches) > 1 {
		return string(matches[1])
	}
	println("ERROR finding submatch", string(e.data))
	return ""
}

func (e *Explosm) getDataFromNet(url string, r *http.Request) ([]byte, error) {
	client := http.DefaultClient
	if r != nil {
		client = urlfetch.Client(appengine.NewContext(r))
	}
	res, err := client.Get(url)
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
