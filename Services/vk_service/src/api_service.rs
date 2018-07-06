use reqwest;
use std::io::Read;
use serde_json;
use serde_json::Value;
use serde_json::map::Map;

use model::json_models::AppSettings;


#[derive(Debug)]
pub struct ApiService<'a> {
    settings: &'a AppSettings,
}

impl<'a> ApiService<'a> {

    pub fn new(settings: &'a AppSettings) -> ApiService {
        ApiService {settings}
    }
    
    pub fn get_news(&self, login: String, password: String) -> Map<String, Value> {
        let token = self.get_token(login, password);

        if token.is_empty() {
            return Map::new();
        }

        let url = format!("https://api.vk.com/method/newsfeed.get?filters=post&access_token={}", token);
        let mut result = match reqwest::get(&url) {
            Ok(resp) => resp,
            Err(_err) => return Map::new(),
        };
        let mut content = String::new();
        result.read_to_string(&mut content).unwrap();


        let json: Value = match serde_json::from_str(&content) {
            Ok(val) => val,
            Err(_err) => return Map::new(),
        };
        let response = match json["response"].as_object() {
            Some(resp) => resp,
            None => return Map::new(),
        };

        return response.clone();
    }

    fn get_token(&self, login: String, password: String) -> String {
        let url = format!(
            "https://oauth.vk.com/token?grant_type=password&client_id={}&client_secret={}&username={}&password={}", 
            self.settings.client_id, self.settings.client_secret, login, password
        );
        let client = reqwest::Client::new();
        let mut result = match client.post(&url).body("").send() {
            Ok(resp) => resp,
            Err(_err) => return String::new(),
        };
        let mut request_result = String::new();
        
        if let Err(_) = result.read_to_string(&mut request_result) {
            return String::new();
        }

        let value: serde_json::Value = match serde_json::from_str(&request_result) {
            Ok(val) => val,
            Err(err) => {
                println!("{:?}", err);
                return String::new();
            },
        };

        return match value.as_object() {
            None => String::new(),
            Some(obj) => match obj["access_token"].as_str() {
                None => String::new(),
                Some(token) => token.to_owned(),
            },
        }
    }

}