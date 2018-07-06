#[derive(Debug, Serialize, Deserialize)]
pub struct AppSettings {
    pub client_id: String,
    pub client_secret: String,
    pub db_host: String,
    pub db_name: String,
    pub db_user: String,
    pub db_password: String,
}