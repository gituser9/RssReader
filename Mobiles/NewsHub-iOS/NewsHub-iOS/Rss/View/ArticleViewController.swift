import UIKit
import WebKit


class ArticleViewController: UIViewController, WKNavigationDelegate {

    @IBOutlet weak var waitIndicator: UIActivityIndicatorView!
    @IBOutlet weak var bodyWebView: WKWebView!
    @IBOutlet weak var linkButton: UIButton!
    
    var articleId = 0
    private var article: Article?
    private var presenter: ArticlePresenter?
    
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        linkButton.isHidden = true
        bodyWebView.isHidden = true
        bodyWebView.navigationDelegate = self

        presenter = ArticlePresenter(view: self)
        presenter?.getArticle(byId: articleId)
    }
    
    
    // MARK: WKNavigationDelegate
    func webView(_ webView: WKWebView, didFinish navigation: WKNavigation!) {
        bodyWebView.isHidden = false
        waitIndicator.stopAnimating()
    }

    
    func showArticle(_ article: Article) {
        self.article = article
        
        linkButton.isHidden = false
        linkButton.setTitle(article.Title, for: .normal)
        
        let html = "<meta name='viewport' content='initial-scale=1.0' /><style>img{max-width:100%;}</style>" + article.Body
        bodyWebView.loadHTMLString(html, baseURL: nil)
    }
    
    
    @IBAction func openArticle(_ sender: UIButton) {
        // todo: open in default browser | open in application (in settings)
        presenter?.openInBrowser(article)
    }
    
}


