import Alamofire


class LoginInteractor {

    func login(login: String, password: String, _ completionHandler: @escaping (User) -> Void) {
        let params: Parameters = [
            "username": login,
            "password": password
        ]
        Alamofire.request(Constant.baseUrl + "auth", method: .post, parameters: params, encoding: JSONEncoding.default).responseJSON { (response) in
            guard let json = response.data else { return }
            do {
                let user = try JSONDecoder().decode(User.self, from: json)

                completionHandler(user)
            } catch {
                // todo: show alert
            }
        }
    }
    
    func getSettings(forUserId userId: Int, _ completionHandler: @escaping (Settings) -> Void) {
        Alamofire.request(Constant.baseUrl + "get-settings?id=\(userId)").responseJSON { (response) in
            guard let json = response.data else { return }
            do {
                let settings = try JSONDecoder().decode(Settings.self, from: json)
                
                completionHandler(settings)
            } catch {
                // todo: show alert
            }            
        }
    }

}
