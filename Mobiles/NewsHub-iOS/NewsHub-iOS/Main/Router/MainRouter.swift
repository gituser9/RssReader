import UIKit


class MainRouter {

    func routeToLogin(_ viewController: MainViewController) {
        let vc = UIStoryboard.init(name: "Login", bundle: nil)
                .instantiateViewController(withIdentifier: "LoginViewController") as! LoginViewController
        viewController.present(vc, animated: true)
    }

    func routeToRss(_ viewController: MainViewController) {
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
