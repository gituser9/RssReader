import UIKit


class MainViewController: UIViewController {

    let presenter = MainPresenter()


    // MARK: Life cycle
    override func viewDidLoad() {
        super.viewDidLoad()
        
        presenter.view = self
        
    }

    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        presenter.checkUserLogin()
    }

}
