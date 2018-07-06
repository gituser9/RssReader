

class VkPresenter {

    var interactor: VkInteractor?
    var view: VkTableViewController?
    
    
    init(view: VkTableViewController) {
        self.view = view
        self.interactor = VkInteractor()
    }
    
    func getPage() {
        let userId = 2
        interactor?.getPage(userId: userId, { [weak self] (vkPage) in
            var viewData = VkPageView()
            
            for item in vkPage.groups {
                viewData.groups[item.gid] = item
            }
            
            viewData.news = vkPage.news
            self?.view?.setPageData(vkPage: viewData)
        })
    }
    
    func getNews(page: Int) {
        interactor?.getNews(page: page, { [weak self] (vkNews) in
            self?.view?.addNews(news: vkNews)
        })
    }
}
