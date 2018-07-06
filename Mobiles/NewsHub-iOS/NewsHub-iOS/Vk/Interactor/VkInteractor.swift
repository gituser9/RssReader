import Alamofire


class VkInteractor {

    func getPage(userId: Int, _ completionHandler: @escaping (VkPage) -> Void) {
        Alamofire.request(Constant.baseUrl + "").responseJSON { (response) in
            guard let json = response.data else { return }
            
            let vkPage = try? JSONDecoder().decode(VkPage.self, from: json)
            
            if vkPage != nil {
                completionHandler(vkPage!)
            }
        }
    }
    
    func getNews(page: Int, _ completionHandler: @escaping ([VkNews]) -> Void) {
        Alamofire.request(Constant.baseUrl + "").responseJSON { (response) in
            // todo: get data from json to common (generic)
            guard let json = response.data else { return }
            
            let vkNews = try? JSONDecoder().decode([VkNews].self, from: json)
            
            if vkNews != nil {
                completionHandler(vkNews!)
            }
        }
    }

}
