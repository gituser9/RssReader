import UIKit


class LoginRouter {

    func routeToRss(_ viewController: LoginViewController) {
        let vc = UIStoryboard.init(name: "Rss", bundle: nil)
                .instantiateViewController(withIdentifier: "FeedListViewController") as! FeedListViewController
        let navigationController = UINavigationController(rootViewController: vc)

        viewController.present(navigationController, animated: true)
    }

    func routeToVk() {

    }

    func routeToTwitter() {

    }

    func routeToSettings() {

    }

}
