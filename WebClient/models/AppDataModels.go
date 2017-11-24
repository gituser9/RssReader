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

type RegistrationData struct {
	User    *Users
	Message string
}

type VkPageData struct {
	News   []VkNews
	Groups []VkGroup
}

type TwitterPageData struct {
	News    []TwitterNewsView
	Sources []TwitterSource
}
type TwitterNewsView struct {
	Id          string
	UserId      int
	SourceId    int
	Text        string
	ExpandedUrl string
	Image       string
}
