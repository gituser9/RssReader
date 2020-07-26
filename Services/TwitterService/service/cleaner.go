package service

import (
	"time"

	"newshub-twitter-service/dao"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Clean() {
	timestamp := time.Now().Add(-((30 * 24) * time.Hour)) // mounth ago
	query := `delete from twitternews where CreatedAt <= ?`

	sources := []dao.TwitterSource{}

	dbExec(func(db *gorm.DB) {
		db.Exec(query, timestamp)
		db.Select("Id").Find(&sources)

		for _, source := range sources {
			cnt := 0
			db.Where(&dao.TwitterNews{SourceId: source.Id}).Count(&cnt)

			if cnt == 0 {
				db.Where(&dao.TwitterSource{Id: source.Id}).Delete(&dao.TwitterSource{})
			}
		}
	})
}
