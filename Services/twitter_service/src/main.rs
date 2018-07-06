#[macro_use]
extern crate serde_derive;
extern crate serde_json;
extern crate base64;
extern crate reqwest;
extern crate hyper;

#[macro_use]
extern crate mysql;

use std::fs::File;
use std::thread;
use std::time;
use std::io::Read;

use model::json_model::AppSettings;
use model::db_models::User;

pub mod model;
pub mod data_provider;
pub mod twitter_service;


fn main() {
    loop {
        let settings = get_settings();
        update(settings);

        thread::sleep(time::Duration::from_secs(1800));
    };
}

fn get_settings() -> AppSettings {
    let mut file = File::open("twitter-settings.json").unwrap();
    let mut content = String::new();
    file.read_to_string(&mut content).unwrap();

    let settings: AppSettings = serde_json::from_str(&content).unwrap();

    return settings;
}

fn update(settings: AppSettings) {
    let users = get_users(&settings);

    if users.is_empty() {
        return;
    }

    let mut provider = data_provider::DataProvider::new(&settings);
    provider.get_token();

    let mut service = twitter_service::TwitterService::new(&settings);

    for user in users {
        let json = provider.get_news_for_user(user.twitter_screen_name);
        service.update(json, user.id);
    }

}

fn get_users(settings: &AppSettings) -> Vec<User> {
    let mut builder = mysql::OptsBuilder::new();
    builder.ip_or_hostname(Some(settings.db_host.clone()))
        .db_name(Some(settings.db_name.clone()))
        .pass(Some(settings.db_password.clone()))
        .user(Some(settings.db_user.clone()))
        .prefer_socket(false);
    let pool = mysql::Pool::new(builder).unwrap();
    let user_ids: Vec<u64> = pool.prep_exec("select UserId from settings where TwitterEnabled = 1", ())
        .map(|result| {
            result.map(|x| x.unwrap()).map(|row| {
                let user_id: u64 = mysql::from_row(row);

                return user_id
            }).collect()
        }).unwrap();
    let users: Vec<User> = pool.prep_exec("select Id, TwitterScreenName from users where id in (:ids) ", (user_ids))
        .map(|result| {
            result.map(|x| x.unwrap()).map(|row| {
                let (id, screen_name) = mysql::from_row(row);

                User {
                    id,
                    twitter_screen_name: screen_name,
                }
            }).collect()
        }).unwrap();

    return users;
}