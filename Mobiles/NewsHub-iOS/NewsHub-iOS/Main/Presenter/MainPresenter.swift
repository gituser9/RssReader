import Foundation


class MainPresenter {

    weak var view: MainViewController?
    var router: MainRouter = MainRouter()

    
    func checkUserLogin() {
        let preferences = UserDefaults.standard

        if preferences.integer(forKey: "userId") == 0 {
            router.routeToLogin(view!)
        } else {
            let rssEnabled = preferences.bool(forKey: Constant.rssEnabledKey)
            let vkEnabled = preferences.bool(forKey: Constant.vkEnabledKey)
            let twitterEnabled = preferences.bool(forKey: Constant.twitterEnabledKey)

            if (rssEnabled) {
                router.routeToRss(view!)
                return
            }
            if (vkEnabled) {
                router.routeToVk()
            }
            if (twitterEnabled) {
                router.routeToTwitter()
            }
        }
    }
}
