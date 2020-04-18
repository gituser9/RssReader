package service

import (
	"log"
	"newshub-vk-service/model"

	"github.com/jinzhu/gorm"
)

func Update() {
	users := getUserIds()

	for _, user := range users {
		if user.VkLogin == "" || user.VkPassword == "" {
			continue
		}

		// получить все группы
		groups := getUserGroups(user.Id)

		// получить последни ид новостей по каждой группе
		lastPostIds := getLastPostId(groups)

		// получить новости из вк
		vkNews, err := getVkNews(user.VkLogin, user.Password)

		if err != nil {
			continue
		}
		for _, news := range vkNews {
			// если такой группы нет - добавить и добавить все новости
			if _, ok := groups[news.GroupId]; !ok {
				addGroup(user.Id)
			}

			// если есть - взять последний ид новости и вноситть только те новости у которых болше
		}

	}
}

func addGroup(userId int64) {

}

func getLastPostId(groups map[int64]model.VkGroup) map[int64]int64 {
	ids := make(map[int64]int64, len(groups))

	dbExec(func(db *gorm.DB) {
		for _, group := range groups {
			news := model.VkNews{}
			db.Model(&model.VkNews{}).
				Where(&model.VkNews{GroupId: group.Gid}).
				Order("Id desc").
				Select("Id").
				First(&news)
			ids[group.Gid] = news.Id
		}
	})

	return ids
}

func getUserGroups(userId int64) map[int64]model.VkGroup {
	groups := []model.VkGroup{}
	result := map[int64]model.VkGroup{}

	dbExec(func(db *gorm.DB) {
		db.Where(&model.VkGroup{UserId: userId}).Find(&groups)
	})
	for _, group := range groups {
		result[group.Gid] = group
	}

	return result
}

func getUserIds() []model.Users {
	users := []model.Users{}

	dbExec(func(db *gorm.DB) {
		ids := []int64{}

		err := db.
			Model(&model.Settings{}).
			Where(&model.Settings{VkNewsEnabled: true}).
			Pluck("UserId", &ids).
			Error
		if err != nil {
			log.Println("get vk users err:", err)
			return
		}

		db.Where("Id in (?)", ids).Find(&users)
	})

	return users
}
