package models

type Configuration struct {
	Driver           string `json:"driver"`
	ConnectionString string `json:"connection_string"`
	UpdateMinutes    int    `json:"update_minutes"`
	AccessToken string `json:"access_token"`
}

type InstagramNewsJson struct {

}

type UserData struct {
    Id string `json:"id"`
    FullName string `json:"full_name"`
    ProfilePicture string `json:"profile_picture"`
}

type UserDataJson struct {
    Data []UserData `json:"data"`
}