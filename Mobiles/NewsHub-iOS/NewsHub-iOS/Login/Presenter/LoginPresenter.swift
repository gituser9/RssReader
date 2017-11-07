import Foundation


class LoginPresenter {
    
    let interactor = LoginInteractor()
    let router = LoginRouter()
    var view: LoginViewController?


    func login(username: String?, password: String?) {
        if username == nil || password == nil {
            // todo: show alert
            return
        }

        interactor.login(login: username!, password: password!) { [weak self] (user)  in
            // save settings
            guard let settings = user.Settings else { return }
            guard let viewController = self?.view else { return }

            self?.saveSettings(settings)

            // routes
            if (settings.RssEnabled) {
                self?.router.routeToRss(viewController)
                return
            }
            if (settings.VkNewsEnabled) {
                self?.router.routeToVk()
            }
            if (settings.TwitterEnabled) {
                self?.router.routeToTwitter()
            }

            self?.router.routeToSettings()
        }
    }

    private func saveSettings(_ settings: Settings) {
        let preferences = UserDefaults.standard

        // todo: consts
        preferences.set(settings.TwitterEnabled, forKey: "twitterEnabled")
        preferences.set(settings.VkNewsEnabled, forKey: "vkEnabled")
        preferences.set(settings.RssEnabled, forKey: "rssEnabled")
        preferences.set(settings.UserId, forKey: "userId")
        preferences.set(settings, forKey: "settings")

        let didSave = preferences.synchronize()

        if !didSave {
            print("Save preferences error")
        }
    }
}
