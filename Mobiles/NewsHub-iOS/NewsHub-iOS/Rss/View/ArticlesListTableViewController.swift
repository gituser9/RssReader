import UIKit


class ArticlesListTableViewController: UITableViewController {
    
    var feedModel: FeedModel?
    private var articles = [ArticleTitle]()
    private var presenter: ArticlesListPresenter?

    override func viewDidLoad() {
        super.viewDidLoad()
        
        presenter = ArticlesListPresenter(view: self)
        presenter?.getArticles(feedId: feedModel?.Feed?.Id ?? 0)
        
        navigationItem.title = feedModel?.Feed?.Name
    }
    

    // MARK: - Table view data source
    
    override func numberOfSections(in tableView: UITableView) -> Int {
        return 1
    }

    override func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return articles.count
    }

    
    override func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "ArticleTitleCell", for: indexPath)
        let article = articles[indexPath.row]
        
        if !article.IsRead {
            // bold font
        }
        
        cell.textLabel?.text = article.Title

        return cell
    }

    
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == "showArticleSegue" {
            guard let nextVc = segue.destination as? ArticleViewController,
                let cell = sender as? UITableViewCell,
                let indexPath = tableView.indexPath(for: cell)
            else {
                return
            }
            let article = articles[indexPath.row]
            nextVc.articleId = article.Id
        }
    }
    
    
    func showArticles(_ articles: [ArticleTitle]) {
        self.articles = articles
        tableView.reloadData()
    }

}
