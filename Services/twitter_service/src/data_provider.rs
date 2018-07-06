use model::json_model::AppSettings;
use base64::encode;
use reqwest;
use reqwest::header::Headers;
use std::io::Read;
use serde_json;


// #[derive(Debug)]
pub struct DataProvider<'a> {
    token: String,
    token_type: String,
    content_type: String,
    settings: &'a AppSettings,
}

impl<'a> DataProvider<'a> {
    pub fn new(app_settings: &'a AppSettings) -> DataProvider {
        DataProvider {
            settings: app_settings,
            token: String::new(),
            token_type: String::new(),
            content_type: String::from("application/x-www-form-urlencoded;charset=UTF-8"),
        }
    }

    pub fn get_token(&mut self) {
        let auth_string = format!("{}:{}", self.settings.client_id, self.settings.client_secret);
        let encoded_auth = encode(&auth_string);

        let mut headers = Headers::new();
        headers.set_raw("Authorization", format!("Basic {}", encoded_auth));
        headers.set_raw("Content-Type", self.content_type.to_owned());

        let params = [("grant_type", "client_credentials")];
        let client = reqwest::Client::new();
        let mut result = client.post("https://api.twitter.com/oauth2/token")
            .headers(headers)
            .form(&params)
            .send()
            .unwrap();

        let mut request_result = String::new();
        result.read_to_string(&mut request_result).unwrap();

        let value: serde_json::Value = serde_json::from_str(&request_result).unwrap();
        self.token = value["access_token"].as_str().unwrap_or("").to_owned();
        self.token_type = value["token_type"].as_str().unwrap_or("").to_owned();
    }

    pub fn get_news_for_user(&self, screen_name: String) -> Vec<serde_json::Value> {
        if self.token_type.is_empty() || self.token.is_empty() {
            return vec![]
        }

        let screen_names = self.get_friend_screen_names(&screen_name);
        let url_template = "https://api.twitter.com/1.1/statuses/user_timeline.json?include_rts=1&exclude_replies=1&screen_name=";
        let client = reqwest::Client::new();
        let mut all_news: Vec<serde_json::Value> = vec![];

        let mut headers = Headers::new();
        headers.set_raw("Authorization", format!("{} {}", self.token_type, self.token));

        for name in screen_names {
            let url = format!("{}{}", url_template, name);
            let mut result = match client.get(&url).headers(headers.clone()).send() {
                Ok(res) => res,
                Err(err) => {
                    println!("{:?}", err);
                    continue;
                },
            };

            let mut request_result = String::new();
            result.read_to_string(&mut request_result);

            let mut value: serde_json::Value = match serde_json::from_str(&request_result) {
                Ok(val) => val,
                Err(err) => {
                    println!("{:?}", err);
                    continue;
                },
            };

            match value.as_array_mut() {
                Some(ref mut arr) => all_news.append(arr),
                None => continue
            };
        }

        return all_news;
    }

    fn get_friend_screen_names(&self, screen_name: &str) -> Vec<String> {
        let url = format!("https://api.twitter.com/1.1/friends/list.json?cursor=-1&screen_name={}&skip_status=true&include_user_entities=false", screen_name);
        let auth_string = format!("{} {}", self.token_type, self.token);

        let mut headers = Headers::new();
        headers.set_raw("Authorization", auth_string);

        let client = reqwest::Client::new();
        let mut result = match client.get(&url).headers(headers).send() {
            Ok(res) => res,
            Err(err) => {
                println!("{:?}", err);
                return vec![];
            },
        };

        let mut request_result = String::new();
        result.read_to_string(&mut request_result);

        let value: serde_json::Value = match serde_json::from_str(&request_result) {
            Ok(val) => val,
            Err(err) => {
                println!("{:?}", err);
                return vec![];
            },
        };
        let users = match value["users"].as_array() {
            Some(arr) => arr,
            None => return vec![],
        };

        let mut names = Vec::new();

        for user in users {
           names.push(user["screen_name"].as_str().unwrap_or("").to_owned());
        }

        return names;
    }
}