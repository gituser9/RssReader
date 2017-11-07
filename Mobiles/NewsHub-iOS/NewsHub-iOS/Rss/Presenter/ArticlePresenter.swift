import UIKit


class ArticlePresenter {
    
    let interactor = ArticleInteractor()
    var view: ArticleViewController
    
    
    init(view: ArticleViewController) {
        self.view = view
    }
    
    func getArticle(byId id: Int) {
        interactor.getArticle(id: id) { [weak self] (article) in
            self?.view.showArticle(article)
        }
    }
    
    func openInBrowser(_ article: Article?) {
        if article == nil {
            return
        }        
        if let url = URL(string: article!.Link) {
            UIApplication.shared.open(url, options: [:], completionHandler: nil)
        } 
    }
    
}
