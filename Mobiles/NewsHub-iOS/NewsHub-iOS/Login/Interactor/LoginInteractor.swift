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

}
