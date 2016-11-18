package models

type Feed struct {
	Rss           Rss
	ArticlesCount int
	ExistUnread   bool
}

type ArticlesJSON struct {
	Articles []RssArticle
	Count    int
}

type AppSettings struct {
	UnreadOnly    bool
	MarkSameRead  bool
	UpdateMinutes int
}
