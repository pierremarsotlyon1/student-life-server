package models

type Xml struct {
	Xml string `xml:"xml"`
	Rss struct {
		Channel struct {
			Language string `xml:"language"`
			Items [] struct {
				Title string `xml:"title"`
				Description string `xml:"description"`
				Author string `xml:"author"`
				PubDate string `xml:"pubDate"`
				Guid string `xml:"guid"`
			} `xml:"item"`
		} `xml:"channel"`
	} `xml:"rss"`
}
