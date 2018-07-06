import UIKit


class VkTableViewController: UITableViewController {
    
    var vkGroups = [Int:VkGroup]()
    var vkNews = [VkNews]() {
        didSet {
            tableView.reloadData()
        }
    }

    override func viewDidLoad() {
        super.viewDidLoad()

        // Uncomment the following line to preserve selection between presentations
        // self.clearsSelectionOnViewWillAppear = false

        // Uncomment the following line to display an Edit button in the navigation bar for this view controller.
        // self.navigationItem.rightBarButtonItem = self.editButtonItem
    }
    
    
    // MARK: - Table view data source

    override func numberOfSections(in tableView: UITableView) -> Int {
        return 1
    }

    override func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return vkNews.count
    }

    
    override func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "VkCell", for: indexPath) as! VkCell
        let item = vkNews[indexPath.row]
        
        cell.groupName.text = vkGroups[item.groupId]?.name
        cell.newBody.loadHTMLString(item.text, baseURL: nil)
        
        // cell.groupImage
        // cell.newsImage

        return cell
    }
    
    

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destinationViewController.
        // Pass the selected object to the new view controller.
    }
    */

    // MARK: - Custom
    func setPageData(vkPage: VkPageView) {
        vkGroups = vkPage.groups
        vkNews = vkPage.news
    }
    
    func addNews(news: [VkNews]) {
        vkNews.append(contentsOf: news)
    }
}
