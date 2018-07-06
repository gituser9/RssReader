#[derive(Serialize, Deserialize)]
pub struct AppSettings {
    pub client_id: String,
    pub client_secret: String,
    pub db_host: String,
    pub db_name: String,
    pub db_user: String,
    pub db_password: String,
}

#[derive(Serialize, Deserialize)]
pub struct TwitterNews {
    pub id: u64,
    pub text: String,
    pub source: String,
    pub image: String,
    pub name: String,
}