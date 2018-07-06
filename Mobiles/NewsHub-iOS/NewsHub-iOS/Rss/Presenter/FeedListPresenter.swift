import Foundation


class FeedListPresenter/* : IFeedListPresenter*/ {
    
    var interactor: FeedListInteractor?
    var view: FeedListViewController?
    
    
    init(view: FeedListViewController) {
        self.view = view
        self.interactor = FeedListInteractor()
    }

    func getFeeds() {
        let userId = getUserId()
        interactor?.getFeeds(userId: userId, { [weak self] (feedModels) in
            self?.view?.showFeeds(feedModels)
        })
    }
    
    
    private func getUserId() -> Int {
        // todo: preferences
        // todo: common
        return 2
    }
}
