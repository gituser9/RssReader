import Alamofire


class FeedListInteractor {
   
    func getFeeds(userId: Int, _ completionHandler: @escaping ([FeedModel]) -> Void) {
        Alamofire.request(Constant.baseUrl + "get-all?id=\(userId)").responseJSON { (response) in
            guard let json = response.data else { return }
            
            let feedModels = try? JSONDecoder().decode([FeedModel].self, from: json)
            
            if feedModels != nil {
                completionHandler(feedModels!)
            }
            
        }
    }
}
