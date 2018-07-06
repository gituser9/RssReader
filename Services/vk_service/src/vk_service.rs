use serde_json::map::Map;
use serde_json::Value;
use mysql;
use std::collections::HashSet;

use model::json_models::AppSettings;
use model::db_models::VkGroup;
use model::db_models::VkNews;
use model::data_models::ImageLink;


pub struct VkService<'a> {
    settings: &'a AppSettings,
}

impl<'a> VkService<'a> {

    pub fn new(settings: &'a AppSettings) -> VkService {
        VkService { settings }
    }

    pub fn update(&self, json: Map<String, Value>, user_id: u64) {
        let _data = self.convert_data(json, user_id);
    }

    fn get_existing_group_ids(&self, user_id: u64) -> Vec<u64> {
        let pool = self.get_pool();
        let mut group_ids: Vec<u64> = pool.prep_exec("select Gid from vkgroups where UserId = :user_id", params!{"user_id" => user_id})
            .map(|result| {
                result.map(|x| x.unwrap()).map(|row| {
                    let group_id: u64 = mysql::from_row(row);

                    return group_id
                }).collect()
            }).unwrap();
        group_ids.sort();

        return group_ids;
    }

    fn get_existing_news_ids(&self, user_id: u64) -> Vec<u64> {
        let pool = self.get_pool();
        let mut news_ids: Vec<u64> = pool.prep_exec("select PostId from vknews where UserId = :user_id", params!{"user_id" => user_id})
            .map(|result| {
                result.map(|x| x.unwrap()).map(|row| {
                    let news_id: u64 = mysql::from_row(row);

                    return news_id
                }).collect()
            }).unwrap();
        news_ids.sort();

        return news_ids;
    }

    fn convert_data(&self, json: Map<String, Value>, user_id: u64) -> Vec<VkNews> {
        let news = match json["items"].as_array() {
            Some(arr) => arr,
            None => return vec![],
        };
        let mut group_ids = self.get_existing_group_ids(user_id);
        let news_ids = self.get_existing_news_ids(user_id);
        let groups = self.get_groups(&json);
        let mut result = vec![];

        println!("{:?}", news.len());

        for item in news {
            let news_json = match item.as_object() {
                Some(obj) => obj,
                None => continue,
            };
            if let Some(i) = news_json["marked_as_ads"].as_u64() {
                if i == 1 {
                    continue;
                }
            }
            let post_id = match news_json["post_id"].as_u64() {
                Some(id) => id,
                None => continue,
            };
            let group_id = match news_json["source_id"].as_i64() {
                Some(id) => (id - (id * 2)) as u64,
                None => continue,
            };
            let mut update_groups = vec![];

            if let Ok(_) = news_ids.binary_search(&post_id) {
                continue;
            }

            match group_ids.binary_search(&group_id) {
                Err(_) => {
                    group_ids.push(group_id);
                    group_ids.sort();
                    self.add_group(group_id, &groups);
                },
                Ok(_) => {
                    update_groups.push(group_id);
                }
            }

            let image_link = self.get_image_link(news_json);
            let entity = VkNews{
                id: post_id,
                user_id: user_id,
                group_id: group_id,
                post_id: post_id,
                text: news_json["text"].as_str().unwrap().to_owned(),
                timestamp: news_json["date"].as_u64().unwrap(),
                image: image_link.image,
                link: image_link.link
            };
            result.push(entity);
        }

        self.update_groups(group_ids, &groups);

        return result;
    }

    fn update_groups(&self, group_ids: Vec<u64>, groups: &Vec<VkGroup>) {
        println!("HHHHHHHHHHHHHHHHHHHHHHHHHHHHH");
        // get db groups
        let pool = self.get_pool();
        let entities: Vec<VkGroup> = pool.prep_exec("select Id, Gid, UserId, Name, LinkedName, Image from vkgroups where Gid in (:ids)", (group_ids))
            .map(|result| {
                result.map(|x| x.unwrap()).map(|row| {
                    let (id, gid, user_id, name, linked_name, image) = mysql::from_row(row);

                    return VkGroup {
                        id: id,
                        gid: gid,
                        user_id: user_id,
                        name: name,
                        linked_name: linked_name,
                        image: image,
                    }
                }).collect()
            }).unwrap();

            println!("{:?}", entities);

        let mut stmt = pool.prepare("update vkgroups set Name = :name, LinkedName = :screen_name, Image = :photo where Gid = :gid").unwrap();

        for entity in entities {
            if let Some(group) = groups.iter().find(|&x| x.gid == entity.gid) {
                let data = params!{
                    "name" => &group.name,
                    "screen_name" => &group.linked_name,
                    "photo" => &group.image,
                };
                // stmt.execute(data).unwrap_or_else(|err| println!("updat{}", group.gid));
                stmt.execute(data).unwrap();
            }
        }
    }

    fn get_groups(&self, json: &Map<String, Value>) -> Vec<VkGroup> {
        let mut groups = vec![];
        let groups_json = match json["groups"].as_array() {
            Some(obj) => obj,
            None => return groups,
        };

        for item in groups_json {
            let mut group = VkGroup::default();

            if let Some(gid) = item["gid"].as_u64() {
                group.gid = gid;
            };
            if let Some(name) = item["name"].as_str() {
                group.name = name.to_owned();
            };
            if let Some(linked_name) = item["screen_name"].as_str() {
                group.linked_name = linked_name.to_owned();
            };
            if let Some(image) = item["photo"].as_str() {
                group.image = image.to_owned();
            };

            groups.push(group);
        }

        return groups;
    }

    fn add_group(&self, group_id: u64, groups: &Vec<VkGroup>) {
        if let Some(group) = groups.iter().find(|&x| x.gid == group_id) {
            let pool = self.get_pool();
            let mut stmt = pool.prepare("insert into vkgroups (gid, UserId, Name, LinkedName, Image) values (:gid, :user_id, :name, :linked_name, :image)").unwrap();
            stmt.execute(params!{
                "gid" => group.gid,
                "user_id" => group.user_id,
                "name" => &group.name,
                "linked_name" => &group.linked_name,
                "image" => &group.image,
            }).unwrap();
        }
    }

    fn get_image_link(&self, json: &Map<String, Value>) -> ImageLink {
        if !json.contains_key("attachment") {
            return ImageLink::default();
        }

        let attachment = match json["attachment"].as_object() {
            Some(obj) => obj,
            None => return ImageLink::default(),
        };
        let post_type = match attachment["type"].as_str() {
            Some(string) => string,
            None => return ImageLink::default(),
        };
        let mut image = String::new();
        let mut link = String::new();

        if post_type == "photo" {
            if attachment.contains_key("photo") {
                let photo = match attachment["photo"].as_object() {
                    Some(obj) => obj,
                    None => return ImageLink::default(),
                };
                if photo.contains_key("src_big") {
                    image = match photo["src_big"].as_str() {
                        Some(string) => string.to_owned(),
                        None => return ImageLink::default(),
                    };
                }
            }
        } else if post_type == "link" {
            let link_object = match attachment["link"].as_object() {
                Some(obj) => obj,
                None => return ImageLink::default(),
            };
            if link_object.contains_key("image_big") {
                image = match link_object["image_big"].as_str() {
                    Some(string) => string.to_owned(),
                    None => return ImageLink::default(),
                };
            } else if link_object.contains_key("image_src") {
                image = match link_object["image_src"].as_str() {
                    Some(string) => string.to_owned(),
                    None => return ImageLink::default(),
                };
            }
            link = match link_object["url"].as_str() {
                Some(string) => string.to_owned(),
                None => return ImageLink::default(),
            };
        } else if post_type == "doc" {
            let doc = match attachment["doc"].as_object() {
                Some(obj) => obj,
                None => return ImageLink::default(),
            };
            if doc.contains_key("ext") {
                let ext = match doc["ext"].as_str() {
                    Some(string) => string,
                    None => return ImageLink::default(),
                };
                if ext == "gif" {
                    image = match doc["url"].as_str() {
                        Some(string) => string.to_owned(),
                        None => return ImageLink::default(),
                    };
                }
            }
        }

        return ImageLink::new(image, link);
    }

    fn get_pool(&self) -> mysql::Pool {
        let mut builder = mysql::OptsBuilder::new();
        builder.ip_or_hostname(Some(self.settings.db_host.clone()))
            .db_name(Some(self.settings.db_name.clone()))
            .pass(Some(self.settings.db_password.clone()))
            .user(Some(self.settings.db_user.clone()))
            .prefer_socket(false);

        return mysql::Pool::new(builder).unwrap();
    }

}