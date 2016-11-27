package models

type Feed struct {
	Feed          Feeds
	ArticlesCount int
	ExistUnread   bool
}

type ArticlesJSON struct {
	Articles []Articles
	Count    int
}

type AppSettings struct {
	UnreadOnly    bool
	MarkSameRead  bool
	UpdateMinutes int
}
