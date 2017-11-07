import Foundation


class MainPresenter {

    weak var view: MainViewController?
    var router: MainRouter = MainRouter()

    
    func checkUserLogin() {
        let preferences = UserDefaults.standard

        if preferences.integer(forKey: "userId") == 0 {
            router.routeToLogin(view!)
        } else {
            guard let settings = preferences.object(forKey: "settings") as? Settings else { return }

            if (settings.RssEnabled) {
                router.routeToRss(view!)
                return
            }
            if (settings.VkNewsEnabled) {
                router.routeToVk()
            }
            if (settings.TwitterEnabled) {
                router.routeToTwitter()
            }
        }
    }
}
