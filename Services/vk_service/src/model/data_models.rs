#[derive(Debug)]
pub struct ImageLink {
    pub image: String,
    pub link: String,
}

impl ImageLink {
    pub fn new(image: String, link: String) -> ImageLink {
        ImageLink { image, link }
    }
}

impl Default for ImageLink {
    fn default() -> ImageLink {
        ImageLink {
            image: String::new(),
            link: String::new()
        }
    }
}