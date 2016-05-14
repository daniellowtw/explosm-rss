package explosm

import "encoding/xml"

func generate(v interface{}) string {
	return `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
` + intermediate(v) + `
</rss>`
}

type CDataTest struct {
	Chardata string `xml:",cdata"`
}

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description CDataTest `xml:"description"`
	Category    []string  `xml:"category"`
	Guid        string    `xml:"guid"`
	PubDate     string    `xml:"pubDate"`
}

type channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Item        []Item `xml:"item"`
}

func intermediate(v interface{}) string {
	s, err := xml.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(s)
}
