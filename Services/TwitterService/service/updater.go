package service

import (
	"log"
	"time"

	"newshub-twitter-service/dao"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Update() {
	users := getUsers()

	for _, user := range users {
		updChan <- user
	}
}

func Close() {
	closeChan <- true
}

func listener() {
	for {
		select {
		case user := <-updChan:
			if user.TwitterScreenName == "" {
				continue
			}

			sources := getSources(user.Id)
			tweetIds := getLastPostId(sources)
			tweets := getNews(user.TwitterScreenName)

			for _, tweet := range tweets {
				if _, ok := sources[tweet.Source.Url]; !ok {
					addSource(user.Id, tweet)
					sources = getSources(user.Id)
					tweetIds = getLastPostId(sources)
				}
				if _, ok := tweetIds[tweet.Id]; !ok {
					addNews(sources, tweet)
					tweetIds[tweet.Id] = true
				}
			}
		case <-closeChan:
			return
		}
	}
}

func addSource(userId int64, data dao.TweetJson) {
	source := data.ToDbSource()
	source.UserId = userId

	dbExec(func(db *gorm.DB) {
		if err := db.Save(&source).Error; err != nil {
			log.Println("save source err:", err)
		}
	})
}

func addNews(sources map[string]dao.TwitterSource, data dao.TweetJson) {
	source := sources[data.Source.Url]
	news := data.ToDbNews()
	news.UserId = source.UserId
	news.SourceId = source.Id
	news.CreatedAt = time.Now().Unix()

	dbExec(func(db *gorm.DB) {
		if err := db.Save(&news).Error; err != nil {
			log.Println("save news err:", err)
		}
	})
}

func getLastPostId(sources map[string]dao.TwitterSource) map[int64]bool {
	ids := make(map[int64]bool)

	dbExec(func(db *gorm.DB) {
		for _, source := range sources {
			news := []dao.TwitterNews{}
			db.Where(&dao.TwitterNews{SourceId: source.Id}).
				Order("Id desc").
				Select("TweetId").
				Find(&news)
			for _, item := range news {
				ids[item.TweetId] = true
			}
		}
	})

	return ids
}

func getSources(userId int64) map[string]dao.TwitterSource {
	sources := []dao.TwitterSource{}
	result := map[string]dao.TwitterSource{}

	dbExec(func(db *gorm.DB) {
		db.Where(&dao.TwitterSource{UserId: userId}).Find(&sources)
	})
	for _, source := range sources {
		result[source.Url] = source
	}

	return result
}

func getUsers() []dao.User {
	users := []dao.User{}

	dbExec(func(db *gorm.DB) {
		ids := []int64{}

		db.Model(&dao.Settings{}).
			Where(&dao.Settings{TwitterEnabled: true}).
			Pluck("UserId", &ids)

		if len(ids) == 0 {
			return
		}

		db.Where("Id in (?)", ids).Find(&users)
	})

	return users
}
