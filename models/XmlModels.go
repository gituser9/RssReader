package models

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

/*==============================================================================
	OPML models
==============================================================================*/

// OPML - struct for OPML file
type OPML struct {
	XMLName  xml.Name      `xml:"opml"`
	Version  float32       `xml:"version,attr"`
	HeadText string        `xml:"head>text"`
	Outlines []OPMLOutline `xml:"body>outline"`
}

// OPMLOutline - RSS description in OPML file
type OPMLOutline struct {
	Title string `xml:"title,attr"`
	URL   string `xml:"xmlUrl,attr"`
	//Version                string `xml:"version, attr"`
	//Description            string `xml:"description,attr"`
	//Type                   string `xml:"type,attr"`
	//ArchiveMode            string `xml:"archiveMode,attr"`
	Text string `xml:"text,attr"`
	//ID                     uint64 `xml:"id"`
	//FetchInterval          int    `xml:"fetchInterval,attr"`
	//MaxArticleAge          int    `xml:"maxArticleAge,attr"`
	//UseCustomFetchInterval bool   `xml:"useCustomFetchInterval,attr"`
}
