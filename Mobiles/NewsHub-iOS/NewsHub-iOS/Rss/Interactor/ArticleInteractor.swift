import Foundation
import Alamofire


class ArticleInteractor {
    
    func getArticle(id: Int, _ completionHandler: @escaping (Article) -> Void) {
        Alamofire.request(Constant.baseUrl + "get-article?id=\(id)").responseJSON { (response) in
            guard let json = response.data else { return }
            
            do {
                let article = try JSONDecoder().decode(Article.self, from: json)
                
                completionHandler(article)
            } catch {
                // todo: show alert
            }            
        }
    }
    
}
