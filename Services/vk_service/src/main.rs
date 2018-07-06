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

use model::json_models::AppSettings;
use model::db_models::User;

pub mod model;
pub mod api_service;
pub mod vk_service;


fn main() {
    loop {
        let settings = get_settings();
        update(settings);

        thread::sleep(time::Duration::from_secs(1800));
    }
}

fn get_settings() -> AppSettings {
    let mut file = File::open("vk-settings.json").unwrap();
    let mut content = String::new();
    file.read_to_string(&mut content).unwrap();

    let settings: AppSettings = serde_json::from_str(&content).unwrap();
    return settings;
}

fn update(settings: AppSettings) {
    let users = get_users(&settings);
    let api_service = api_service::ApiService::new(&settings);
    let vk_service = vk_service::VkService::new(&settings);

    for user in users {
        let news_json = api_service.get_news(user.vk_login, user.vk_password);
        vk_service.update(news_json, user.id);
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
    let user_ids: Vec<u64> = pool.prep_exec("select UserId from settings where VkNewsEnabled = 1", ())
        .map(|result| {
            result.map(|x| x.unwrap()).map(|row| {
                let user_id: u64 = mysql::from_row(row);

                return user_id
            }).collect()
        }).unwrap();
    let users: Vec<User> = pool.prep_exec("select Id, VkLogin, VkPassword from users where id in (:ids) ", user_ids)
        .map(|result| {
            result.map(|x| x.unwrap()).map(|row| {
                let (id, vk_login, vk_password) = mysql::from_row(row);

                User {
                    id,
                    vk_login,
                    vk_password,
                }
            }).collect()
        }).unwrap();

    return users;
}