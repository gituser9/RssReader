import Foundation


class ArticlesListPresenter {
    
    let interactor = ArticlesListInteractor()
    var view: ArticlesListTableViewController
    
    init(view: ArticlesListTableViewController) {
        self.view = view
    }
    
    func getArticles(feedId: Int) {
        let userId = getUserId()
        interactor.getArticles(feedId: feedId, userId: userId) { [weak self] (articles) in
            var titles = [ArticleTitle]()
            
            for title in articles.Articles {
                titles.append(title)
            }
            
            self?.view.showArticles(titles)
        }
    }
    
    func getUserId() -> Int {
        return 2
    }
    
}
