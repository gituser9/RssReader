#[derive(Debug, PartialEq, Eq)]
pub struct User {
    pub id: u64,
    pub vk_login: String,
    pub vk_password: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct VkGroup {
    pub id: u64,
    pub gid: u64,
    pub user_id: u64,
    pub name: String,
    pub linked_name: String,
    pub image: String,
}

impl Default for VkGroup {
    fn default() -> VkGroup {
        VkGroup{
            id: 0,
            gid: 0,
            user_id: 0,
            name: String::new(),
            linked_name: String::new(),
            image: String::new(),
        }
    }
}

pub struct VkNews {
    pub id: u64,
    pub user_id: u64,
    pub group_id: u64,
    pub post_id: u64,
    pub timestamp: u64,
    pub text: String,
    pub image: String,
    pub link: String,
}