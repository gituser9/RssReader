package com.newshub.newshub_android.rss.view

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.newshub.newshub_android.R
import com.newshub.newshub_android.rss.model.Article
import com.newshub.newshub_android.rss.presenter.ArticlePresenter
import kotlinx.android.synthetic.main.fragment_article.*


class ArticleFragment : BaseRssFragment() {

    var articleId: Int? = null
    private var link: String? = null
    private val presenter = ArticlePresenter(this)


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        if (articleId == null) {
            shortToast("Get article error")
        }

        presenter.getArticle(articleId!!)
    }

    override fun onCreateView(inflater: LayoutInflater?, container: ViewGroup?, savedInstanceState: Bundle?): View? {
        return inflater!!.inflate(R.layout.fragment_article, container, false)
    }

    override fun onViewCreated(view: View?, savedInstanceState: Bundle?) {
        btnLink.setOnClickListener {
            if (link != null) {
                val intent = Intent(Intent.ACTION_VIEW, Uri.parse(link))
                startActivity(intent)
            }
        }
    }

    fun showArticle(article: Article) {
        btnLink.text = article.title
        btnLink.visibility = View.VISIBLE

        val data = "<style>img { max-width:100%; }</style>" + article.body      // FIXME: костыль злоебучий
        webView.loadData(data, "text/html; charset=utf-8", "UTF-8")

        link = article.link
    }

}
