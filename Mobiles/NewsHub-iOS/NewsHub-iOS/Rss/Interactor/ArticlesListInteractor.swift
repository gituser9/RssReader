import Foundation
import Alamofire


class ArticlesListInteractor {
    
    func getArticles(feedId: Int, userId: Int, _ completionHandler: @escaping (Articles) -> Void) {
        Alamofire.request(Constant.baseUrl + "get-articles?id=\(feedId)&userId=\(userId)").responseJSON { (response) in
            guard let json = response.data else { return }
            
            do {
                let articles = try JSONDecoder().decode(Articles.self, from: json)
                
                if articles.Count > 0 {                    
                    completionHandler(articles)
                }
            } catch {
                // show alert
            }
            
        }
    }
    
}
