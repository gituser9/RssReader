use mysql as my;

#[derive(Debug, PartialEq, Eq)]
pub struct TwitterNews {
    pub id: u64,
    pub user_id: u64,
    pub source_id: u64,
    pub text: String,
    pub image: String,
    pub expanded_url: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct TwitterSource {
    pub id: u64,
    pub user_id: u64,
    pub name: String,
    pub screen_name: String,
    pub url: String,
    pub image: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct User {
    pub id: u64,
    pub twitter_screen_name: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Settings {
    pub id: u64,
    pub user_id: u64,
    pub twitter_news_enabled: bool,
}