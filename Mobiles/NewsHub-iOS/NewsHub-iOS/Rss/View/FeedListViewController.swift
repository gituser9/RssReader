import UIKit

class FeedListViewController: UIViewController, UITableViewDataSource, UITableViewDelegate {

    @IBOutlet weak var tableView: UITableView!
    
    var presenter: FeedListPresenter?
    var feedModels = [FeedModel]()
    var currentModel: FeedModel?
    

    // MARK: Life cycle
    override func viewDidLoad() {
        super.viewDidLoad()
        
        tableView.delegate = self
        tableView.dataSource = self
        
        presenter = FeedListPresenter(view: self)
        presenter?.getFeeds()
    }
    
    override func awakeFromNib() {
        super.awakeFromNib()
        // fixme
//        FeedListAssembly.instance().inject(into: self)
    }
    
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        
        if !feedModels.isEmpty {
            // return from articles
            tableView.reloadData()
        }
    }
    
    
    func showFeeds(_ feeds: [FeedModel]) {
        feedModels = feeds
        tableView.reloadData()
    }

    // MARK: UITableViewDataSource
    func numberOfSections(in tableView: UITableView) -> Int {
        return 1
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return feedModels.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "FeedCell", for: indexPath)
        let feedModel = feedModels[indexPath.row]

        if feedModel.ExistUnread {
            // todo: bold
            cell.textLabel?.text = "\(feedModel.Feed?.Name ?? "") (\(feedModel.ArticlesCount))"
        } else {
            cell.textLabel?.text = feedModel.Feed?.Name
        }

        return cell
    }

    // MARK: UITableViewDelegate
    /*func tableView(_ tableView: UITableView, didDeselectRowAt indexPath: IndexPath) {
        currentModel = feedModels[indexPath.row]
        
        self.performSegue(withIdentifier: <#T##String#>, sender: <#T##Any?#>)
    }*/
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == "showArticlesListSegue" {
            guard let nextVc = segue.destination as? ArticlesListTableViewController,
                  let cell = sender as? UITableViewCell,
                  let indexPath = tableView.indexPath(for: cell)
            else {
                return                
            }
            nextVc.feedModel = feedModels[indexPath.row]
        }
    }

}
