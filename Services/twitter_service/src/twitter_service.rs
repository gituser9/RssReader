use serde_json::map::Map;
use serde_json::Value;
use mysql;

use model::json_model::AppSettings;
use model::db_models::{TwitterNews, TwitterSource};


pub struct TwitterService<'a> {
    user_id: u64,
    settings: &'a AppSettings,
    existing_news_ids: Vec<u64>,
    existing_source_ids: Vec<u64>,
}


impl<'a> TwitterService<'a> {
    pub fn new(settings: &'a AppSettings) -> TwitterService {
        TwitterService {
            user_id: 0,
            settings,
            existing_news_ids: vec![],
            existing_source_ids: vec![],
        }
    }

    pub fn update(&mut self, news_json: Vec<Value>, user_id: u64) {
        self.user_id = user_id;
        self.prepare_data();
        self.save_news(news_json);
    }

    fn prepare_data(&mut self) {
        let mut builder = mysql::OptsBuilder::new();
        builder.ip_or_hostname(Some(self.settings.db_host.clone()))
            .db_name(Some(self.settings.db_name.clone()))
            .pass(Some(self.settings.db_password.clone()))
            .user(Some(self.settings.db_user.clone()))
            .prefer_socket(false);
        let pool = mysql::Pool::new(builder).unwrap();
        self.existing_news_ids = pool.prep_exec("select Id from twitternews where UserId = :user_id", params!{"user_id" => self.user_id})
            .map(|result| {
                result.map(|x| x.unwrap()).map(|row| {
                    let user_id: u64 = mysql::from_row(row);

                    return user_id
                }).collect()
            }).unwrap();
        self.existing_source_ids = pool.prep_exec("select Id from twittersource where UserId = :user_id", params!{"user_id" => self.user_id})
            .map(|result| {
                result.map(|x| x.unwrap()).map(|row| {
                    let user_id: u64 = mysql::from_row(row);

                    return user_id
                }).collect()
            }).unwrap();
        self.existing_news_ids.sort();
    }

    fn save_news(&mut self, news_json: Vec<Value>) {
        let news = self.get_news(&news_json);
        let mut builder = mysql::OptsBuilder::new();
        builder.ip_or_hostname(Some(self.settings.db_host.clone()))
            .db_name(Some(self.settings.db_name.clone()))
            .pass(Some(self.settings.db_password.clone()))
            .user(Some(self.settings.db_user.clone()))
            .prefer_socket(false);
        let pool = match mysql::Pool::new(builder) {
            Ok(p) => p,
            Err(err) => {
                println!("get connection pool error {:?}", err);
                return;
            },
        };

        for mut stmt in pool.prepare(
            "insert into twitternews (Id, UserId, SourceId, Text, ExpandedUrl, Image) values (:id, :user_id, :source_id, :text, :expanded_url, :image)"
        ).into_iter() {
            for item in news.iter() {
                match stmt.execute(params!{
                    "id" => item.id,
                    "user_id" => item.user_id,
                    "source_id" => item.source_id,
                    "text" => item.text.to_owned(),
                    "expanded_url" => item.expanded_url.to_owned(),
                    "image" => item.image.to_owned(),
                }) {
                    Err(err) => println!("insert into twitternews error {:?}", err),
                    Ok(_) => (),
                };
            }
        }
        for mut stmt in pool.prepare(
            "insert into twittersource (Id, UserId, Name, ScreenName, Url, Image) values (:id, :user_id, :name, :screen_name, :url, :image)"
        ).into_iter() {
            for source in self.get_sources(&news_json) {
                match stmt.execute(params!{
                    "id" => source.id,
                    "user_id" => source.user_id,
                    "name" => source.name.to_owned(),
                    "screen_name" => source.screen_name.to_owned(),
                    "url" => source.url.to_owned(),
                    "image" => source.image.to_owned(),
                }) {
                    Err(err) => println!("insert into twittersource error {:?}", err),
                    Ok(_) => (),
                };
            }
        }
    }

    fn get_news(&self, news_json: &Vec<Value>) -> Vec<TwitterNews> {
        let mut news = Vec::with_capacity(news_json.len());
        let user_id = self.user_id;

        for item in news_json {
            let news_object = match item.as_object() {
                Some(obj) => obj,
                None => continue,
            };
            let id = match news_object["id"].as_u64() {
                Some(id) => id,
                None => continue,
            };

            match self.existing_news_ids.binary_search(&id) {
                Err(_) => println!(),
                Ok(_) => continue
            }

            let source_id = match news_object["user"].as_object() {
                None => continue,
                Some(user) => match user["id"].as_u64() {
                    None => continue,
                    Some(id) => id,
                }
            };
            let text = match news_object["text"].as_str() {
                Some(txt) => txt,
                None => continue,
            };
            let mut entity = TwitterNews {
                id,
                user_id,
                source_id,
                text: String::from(text),
                image: String::new(),
                expanded_url: String::new(),
            };
            entity.expanded_url = self.get_expanded_url(news_object);
            entity.image = self.get_image(news_object);

            news.push(entity);
        }

        return news;
    }

    fn get_expanded_url(&self, json: &Map<String, Value>) -> String {
        if !json.contains_key("entities") {
            return String::new();
        }

        let entities = match json["entities"].as_object() {
            Some(obj) => obj,
            None => return String::new(),
        };

        if !entities.contains_key("urls") {
            return String::new();
        }

        let urls = match entities["urls"].as_array() {
            Some(arr) => {
                if arr.is_empty() {
                    return String::new();
                }
                arr
            },
            None => return String::new(),
        };

        let url = match urls.first() {
            None => return String::new(),
            Some(obj) => match obj.as_object() {
                None => return String::new(),
                Some(data) => data,
            },
        };

        return url["expanded_url"].as_str().unwrap_or("").to_owned();
    }

    fn get_image(&self, json: &Map<String, Value>) -> String {
        if !json.contains_key("entities") {
            return String::new();
        }

        let entities = match json["entities"].as_object() {
            Some(obj) => obj,
            None => return String::new(),
        };

        if !entities.contains_key("media") {
            return String::new();
        }

        let media = match entities["media"].as_array() {
            None => return String::new(),
            Some(arr) => match arr.first() {
                None => return String::new(),
                Some(val) => match val.as_object() {
                    None => return String::new(),
                    Some(obj) => obj,
                },
            },
        };

        if !media["type"].as_str().unwrap_or("").eq("photo") {
            return String::new();
        }
        if media.contains_key("media_url_https") {
            return media["media_url_https"].as_str().unwrap_or("").to_owned();
        }
        if media.contains_key("media_url_http") {
            return media["media_url_http"].as_str().unwrap_or("").to_owned();
        }

        return String::new();
    }

    fn get_sources(&mut self, json: &Vec<Value>) -> Vec<TwitterSource> {
        let mut sources: Vec<TwitterSource> = vec![];

        for item in json {
            let news_object = match item.as_object() {
                Some(obj) => obj,
                None => continue,
            };
            let source_id = match news_object["user"].as_object() {
                None => continue,
                Some(user) => match user["id"].as_u64() {
                    None => continue,
                    Some(id) => id,
                },
            };

            match self.existing_source_ids.binary_search(&source_id) {
                Err(_) => println!(),
                Ok(_) => continue
            }

            let user = match news_object["user"].as_object() {
                Some(obj) => obj,
                None => continue,
            };
            let name = match user["name"].as_str() {
                Some(string) => string.to_owned(),
                None => continue,
            };
            let screen_name = match user["screen_name"].as_str() {
                Some(string) => string.to_owned(),
                None => continue,
            };
            let mut url = String::new();

            if user.contains_key("url") {
                url = user["url"].as_str().unwrap_or("").to_owned();
            }

            let mut entity = TwitterSource {
                id: source_id,
                user_id: self.user_id,
                name,
                screen_name,
                url,
                image: user["profile_image_url"].as_str().unwrap_or("").to_owned()
            };

            sources.push(entity);
            self.existing_source_ids.push(source_id);
            self.existing_source_ids.sort();
        }

        return sources;
    }
}