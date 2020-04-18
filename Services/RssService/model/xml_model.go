package model

import "encoding/xml"

// XMLFeed - struct for RSS XML
type XMLFeed struct {
	XMLName     xml.Name     `xml:"rss"`
	Version     string       `xml:"version,attr"`
	RssName     string       `xml:"channel>title"`
	RssURL      string       `xml:"channel>link"`
	Description string       `xml:"channel>description"`
	Articles    []XMLArticle `xml:"channel>item"`
}

// XMLArticle - article in RSS XML
type XMLArticle struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Date        string `xml:"pubDate"`
	IsRead      bool
}